package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(alphabet string, length int) string {
	var password string
	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		num := rand.Intn(len([]rune(alphabet)))
		password += string([]rune(alphabet)[num])
	}

	return password
}

func bruteForce(alphabet string, password string, size int, currPassword *string) {
	if *currPassword == password {
		return
	}
	if len([]rune(*currPassword)) == size {
		return
	}

	for _, value := range []rune(alphabet) {

		*currPassword += string(value)

		bruteForce(alphabet, password, size, currPassword)
		if *currPassword == password {
			return
		}

		*currPassword = string([]rune(*currPassword)[:len([]rune(*currPassword))-1])

	}

}

func main() {

	var alphabet string
	alphabet = "абвгдеёжзийклмнопрстуфхцшщъыьэюя"

	var length int

	fmt.Print("Password length: ")

	for _, err := fmt.Scan(&length); err != nil; _, err = fmt.Scan(&length) {
		fmt.Println("Invalid input")
	}

	password := generatePassword(alphabet, length)

	fmt.Println("Password:", password)

	var currPassword string
	var size int = 1
	var crackTime int64

	crackTime = time.Now().UnixMicro()

	for {
		bruteForce(alphabet, password, size, &currPassword) // currPassword is the result of hacking
		if currPassword == password {
			break
		}
		size++
	}

	crackTime = time.Now().UnixMicro() - crackTime

	fmt.Println("Found:", currPassword)
	fmt.Println("crackTime:", crackTime, "Microseconds")
}
