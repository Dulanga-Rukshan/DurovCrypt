package main

import (
	ess "DurovCrypt/essentials"
	"os"

	"fmt"

	// "os"
	"strings"
)

const (
	encryptFunctionCall = "Encrypt"
	decryptFunctionCall = "Decrypt"

	//ERRORTYPE
	aerrorType    = "OASKERR: "      //ask error
	ferrorType    = "OFILEERR: "     //file path ask error
	perrorType    = "OPASKERR: "     //password ask error
	oefaerrorType = "OPENCRYFAERR: " //operation encrypt faliure error
	odfaerrorType = "OPDECRYFAERR: " //operation decrypt faliure error
)

func operationToPerform(operation string) (string, error) {

	//getting file path for decrypt
	fileName, err := ess.FilePathInput(operation)
	if err != nil {
		if err.Error() == "operation cancelled by user" {
			fmt.Println("EXIT: Operation cancelled by user")
			os.Exit(0)
		}
		ess.MainErr("", err)
	}

	switch strings.ToUpper(operation) {
	case "ENCRYPT", "E":

		//ask user for password to encrypt
		password, err := ess.PasswordAskInput(encryptFunctionCall, fileName)
		ess.MainErr("", err)

		//open the file and read the data and assign data to variable
		fileData, _, _, err := ess.FileRead(fileName, operation)
		ess.MainErr("", err)

		//encrypting the data
		cipherText, derivedSalt, derivedNonce, err := ess.Encrypt(password, fileData)
		ess.MainErr("", err)

		//write the ciphertext data to file
		successMsg, err8 := ess.FileWrite(cipherText, fileName, derivedSalt, derivedNonce, operation)
		return successMsg,
			err8

	case "DECRYPT", "D":

		//ask user for password to encrypt
		successMsg, err := ess.PasswordAskInput(decryptFunctionCall, fileName)
		if err != nil {
			ess.MainErr("", err)
			os.Exit(0) //exiting the program after maximum attempts reach
		}
		return successMsg, nil
	}

	return "",
		nil
}
func main() {
	fmt.Println(ess.WelcomeMsg())
	Operation, err := ess.DefaultAskInput()
	ess.MainErr(aerrorType, err)

	switch strings.ToUpper(Operation) {

	//if user input 'e' for encrypt
	case "E", "ENCRYPT":
		successMsg, err := operationToPerform(encryptFunctionCall)
		ess.MainErr(oefaerrorType, err)
		fmt.Println(successMsg)

	case "D", "DECRYPT":
		successMsg, err := operationToPerform(decryptFunctionCall)
		ess.MainErr(odfaerrorType, err)
		fmt.Println(successMsg)
	case "H", "HELP":
		helpMsg := ess.ShowHelp()
		fmt.Println(helpMsg)
	}
}
