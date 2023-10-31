package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(length int) string {
	var password string
	var alphabet = "абвгдеёжзийклмнопрстуфхцшщъыьэюя"
	var letterNum int

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		letterNum = rand.Intn(len([]rune(alphabet)))
		password += string([]rune(alphabet)[letterNum])
	}

	return password
}

func getString(matrix [][]string) string {
	var currentString string

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			currentString += matrix[i][j]
		}
	}
	return currentString
}

func addSymbols(maxRowLen int, matrix *[][]string) {

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < len(*matrix); i++ {
		if len((*matrix)[i]) < maxRowLen {
			for j := 0; j < maxRowLen-len((*matrix)[i]); j++ {
				(*matrix)[i] = append((*matrix)[i], string("ю"))
			}
		}
	}
}

func tableCipher(password string, M int) (string, int) { // M is the number of rows of the table for encrypting
	var cipherMatrix [][]string
	var numRows int
	numRows = M
	var maxRowLen = 0

	for i := 0; i < numRows; i++ {
		var currentRow []string

		for j := i; j < len([]rune(password)); j += numRows {
			currentRow = append(currentRow, string([]rune(password)[j]))
		}
		if len(currentRow) > maxRowLen {
			maxRowLen = len(currentRow)
		}
		cipherMatrix = append(cipherMatrix, currentRow)
	}
	addSymbols(maxRowLen, &cipherMatrix)

	return getString(cipherMatrix), len(cipherMatrix[0])
}

func decrypting(M int, N int, cipher string) string {
	var password string

	for i := 0; len([]rune(password)) != len([]rune(cipher)); i++ {
		for j := i; j < len([]rune(cipher)); j += N {
			password += string([]rune(cipher)[j])
		}
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
	var password string
	var passwordLength int

	fmt.Println("Enter password length")
	for _, err := fmt.Scan(&passwordLength); err != nil; _, err = fmt.Scan(&passwordLength) {
		fmt.Println("Enter valid number:")
	}

	password = generatePassword(passwordLength)

	fmt.Println("Generated password:\n" + password)

	var cipher string
	var M, N int
	M = passwordLength - 1
	cipher, N = tableCipher(password, M)

	fmt.Println("Encrypted password:\n" + cipher)
	decryptedPassword := decrypting(M, N, cipher)
	fmt.Println("Decrypted password:\n" + decryptedPassword)
	fmt.Println("Table size: ", M, "x", N)

	var alphabet = "абвгдеёжзийклмнопрстуфхцшщъыьэюя"
	var currPassword string
	var size = 1 // var size is a len of currPassword

	start := time.Now().UnixNano()

	var currentCipher string
	var keyM, keyN int

	for {
		bruteForce(alphabet, password, size, &currPassword)

		keyM = 1
		for currentCipher != cipher {
			currentCipher, keyN = tableCipher(currPassword, keyM)

			if keyM == len([]rune(cipher)) {
				break
			}

			keyM++
		}

		if currentCipher == cipher {
			break
		}
		size++
	}

	crackingTime := time.Now().UnixNano() - start

	fmt.Println("Key size:", keyM-1, "x", keyN)
	fmt.Println("Current password: " + currPassword)
	fmt.Println("Cracking time = ", crackingTime)

}
