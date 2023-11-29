package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	rn "math/rand"
	"os"
	"time"
)

func getPrime() *big.Int {
	a, err := rand.Prime(rand.Reader, 64)
	if err != nil {
		panic(err.Error())
	}

	return a
}

func getE() *big.Int {
	source := rn.NewSource(time.Now().UnixNano())
	random := rn.New(source)
	size := random.Int31n(4) + 2
	a, err := rand.Prime(rand.Reader, int(size))
	if err != nil {
		panic(err.Error())
	}

	return a
}

type WriteJsonPublic struct {
	PublicKey1 *big.Int `json:"pkey1"`
	PublicKey2 *big.Int `json:"pkey2"`
}

type WriteJsonPrivate struct {
	PrivateKey1 *big.Int `json:"prkey1"`
	PrivateKey2 *big.Int `json:"prkey2"`
}

type Message struct {
	Message string `json:"message"`
}

type CiMessage struct {
	Message []*big.Int `json:"message"`
}

func cipher() {
	var pub WriteJsonPublic
	arr, _ := os.ReadFile("data_public.json")

	err := json.Unmarshal(arr, &pub)
	if err != nil {
		fmt.Println(err.Error())
	}

	var message Message
	arr, _ = os.ReadFile("message.json")

	err = json.Unmarshal(arr, &message)
	if err != nil {
		fmt.Println(err.Error())
	}

	var ciMessage []*big.Int

	for _, i := range message.Message {
		curr := new(big.Int)
		ciMessage = append(ciMessage, curr.Mod(curr.Exp(big.NewInt(int64(i)), pub.PublicKey1, nil), pub.PublicKey2))
	}
	var obj CiMessage
	obj.Message = ciMessage

	arr, err = json.MarshalIndent(obj, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("cimessage.json", arr, 'w')
	if err != nil {
		fmt.Println(err.Error())
	}
}

func decipher() {
	var pri WriteJsonPrivate
	arr, _ := os.ReadFile("data_private.json")

	err := json.Unmarshal(arr, &pri)
	if err != nil {
		fmt.Println(err.Error())
	}

	var ciMessage CiMessage

	arr, _ = os.ReadFile("cimessage.json")

	err = json.Unmarshal(arr, &ciMessage)

	// var word []string
	for _, i := range ciMessage.Message {
		curr := new(big.Int)
		fmt.Print(string(rune(curr.Mod(curr.Exp(i, pri.PrivateKey1, nil), pri.PrivateKey2).Int64())), " ")
		fmt.Println(curr.Mod(curr.Exp(i, pri.PrivateKey1, nil), pri.PrivateKey2))
		//word = append(word, string(rune(curr.Mod(curr.Exp(i, pri.PrivateKey1, nil), pri.PrivateKey2).Int64())))
	}

	// return word
}

func main() {
	q := getPrime()
	p := getPrime()

	n1 := new(big.Int)
	n2 := new(big.Int)

	n1.Add(q, big.NewInt(-1))
	n2.Add(p, big.NewInt(-1))

	n := new(big.Int)
	n.Mul(n1, n2)

	e := getE()

	d := new(big.Int).ModInverse(e, n)

	fmt.Println("Private key: ", d, ",", n)
	fmt.Println("\n\nPublic key: ", e, ", ", n, "\n")

	structJsonPub := WriteJsonPublic{PublicKey1: e, PublicKey2: n}
	structJsonPri := WriteJsonPrivate{PrivateKey1: d, PrivateKey2: n}

	arr, _ := json.MarshalIndent(structJsonPub, "", "\t")

	err := os.WriteFile("data_public.json", arr, 'w')
	if err != nil {
		fmt.Println(err.Error())
	}

	arr, _ = json.MarshalIndent(structJsonPri, "", "\t")

	err = os.WriteFile("data_private.json", arr, 'w')
	if err != nil {
		fmt.Println(err.Error())
	}

	cipher()
	decipher()
}
