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

func operationToPerform(operation string) (string, error) {
	var (
		result     string
		successMsg string
		err        error
		err8       error
	)
	//getting file path for decrypt
	fileName, err2 := ess.FilePathInput(operation)
	ess.MainErr(err2)

	//ask user for password to encrypt
	password, err3 := ess.PasswordAskInput(operation)
	ess.MainErr(err3)

	switch strings.ToUpper(operation) {
	case "ENCRYPT", "E":

		//open the file and read the data and assign data to variable
		fileData, _, err6 := ess.FileRead(fileName, operation)
		ess.MainErr(err6)

		//encrypting the data
		cipherText, derivedSalt, err := ess.Encrypt(password, operation, string(fileData))
		ess.MainErr(err)

		//write the ciphertext data to file
		successMsg, err8 := ess.FileWrite([]byte(cipherText), fileName, []byte(derivedSalt), operation)
		return successMsg, err8

	case "DECRYPT", "D":
		//open the file and read the data and assign data to variable
		fileData, derivedSalt, err6 := ess.FileRead(fileName, operation)
		ess.MainErr(err6)

		fmt.Print(string(derivedSalt))

		//key generation
		derivedKey, _, err5 := ess.KeyGen(password, operation, derivedSalt)
		ess.MainErr(err5)

		//encrypting the data
		result, err = ess.Decrypt(derivedKey, string(fileData))

		//write the ciphertext data to file
		successMsg, err8 := ess.FileWrite([]byte(result), fileName, nil, operation)
		return successMsg, err8
	}

	ess.MainErr(err)
	return successMsg, err8
}
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
		successMsg, err := operationToPerform(encryptFunctionCall)
		ess.MainErr(err)
		fmt.Println(successMsg)

	case strings.ToUpper(Operation) == "D" || strings.ToUpper(Operation) == "DECRYPT":
		successMsg, err := operationToPerform(decryptFunctionCall)
		ess.MainErr(err)
		fmt.Println(successMsg)
	}
}
