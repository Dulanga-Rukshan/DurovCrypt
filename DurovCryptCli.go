package main

import (
	ess "DurovCrypt/essentials"
	"fmt"

	// "os"
	"runtime"
	"strings"
)

func main() {

	//password policy
	Policy := ess.PasswordPolicyCheck{
		MinLength:           6,
		MaxLength:           64,
		RequireUpper:        true,
		RequireLower:        true,
		RequireSymbol:       true,
		RequireNumber:       true,
		DenyWhiteSpace:      true,
		AllowedSpecialChars: "{}()!@#$%^&*+_<>?:';,.][-|\"\\/",
	}

	//Welcome message

	// args := os.Args

	// if len(args) > 1{
	// 	if
	// }
	//ask user for what to do
	Operation, err1 := ess.InputPrompt("What operation do you wanna perform: ")
	if err1 != nil {
		fmt.Printf("ERR: Error getting input prompt. \n\n%v", err1)

	}

	switch {

	//if user input 'e' for encrypt
	case strings.ToUpper(Operation) == "E":

		//getting file path for encrypt
		fileName, err2 := ess.FilePathInput()
		if err2 != nil {
			fmt.Printf("\nERROR: %v\n", err2) // Prints the error from IsValidFileName
			return
		}

		//ask user for password to encrypt
		password, err3 := ess.InputPrompt("Enter a Password for encrypt: ")
		if err3 != nil {
			fmt.Printf("\nERROR: %v\n", err3)
			return
		}
		err4 := ess.PasswordChecker(password, Policy)
		if err4 != nil {
			fmt.Printf("\nERROR: %v\n", err4)
			return
		}

		//key generation
		salt := ess.DataKey()
		NewKey := &ess.Aragon2Key{
			Password:  []byte(password),
			Salt:      salt,
			Iteration: uint32(3),
			MemSize:   uint32(64 * 1024),
			Threads:   uint8(runtime.NumCPU()),
			KeyLength: 32,
		}

		derivedKey, err6 := NewKey.Generate()
		if err6 != nil {
			fmt.Printf("\nERROR: %v\n", err6)
			return
		}

		//open the file and read it
		fileData, err7 := ess.FileRead(fileName)
		if err7 != nil {
			fmt.Printf("\nERROR: %v\n", err7)
			return
		}

		//encrypting the data
		encryptedData, err8 := ess.Encrypt(derivedKey, string(fileData))
		if err8 != nil {
			fmt.Printf("\nERR: %v\n", err8)
			return
		}
		fmt.Print(encryptedData)
		//
	}

}
