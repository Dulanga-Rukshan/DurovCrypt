package main

import (
	ess "DurovCrypt/essentials"
	"fmt"

	// "os"
	"strings"
)

const (
	encryptFunctionCall = "Encrypt"
	decryptFunctionCall = "Decrypt"
)

func main() {

	//Welcome message

	//args := os.Args

	// if len(args) > 1{
	// 	if
	// }
	//ask user for what to do
	Operation, err1 := ess.DefaultAskInput("What operation do you wanna perform: ")
	ess.MainErr(err1)

	switch {

	//if user input 'e' for encrypt
	case strings.ToUpper(Operation) == "E" || strings.ToUpper(Operation) == "ENCRYPT":
		//getting file path for encrypt
		fileName, err2 := ess.FilePathInput(encryptFunctionCall)
		ess.MainErr(err2)

		//ask user for password to encrypt
		password, err3 := ess.PasswordAskInput(encryptFunctionCall)
		ess.MainErr(err3)

		//key generation
		derivedKey, err5 := ess.KeyGen(password)
		ess.MainErr(err5)

		//open the file and read the data and assign data to variable
		fileData, err6 := ess.FileRead(fileName)
		ess.MainErr(err6)

		//encrypting the data
		encryptedData, err7 := ess.Encrypt(derivedKey, string(fileData))
		ess.MainErr(err7)

		//write the ciphertext data to file
		successMsg, err8 := ess.FileWrite([]byte(encryptedData), fileName)
		ess.MainErr(err8)

		fmt.Printf("File Encrypted. %v.", successMsg)

		// case strings.ToUpper(Operation) == "D" || strings.ToUpper(Operation) == "DECRYPT":
		// 	OperationForEnAndDe(decryptFunctionCall)
	}
}
