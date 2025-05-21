package DurovCrypt

import (
	"errors"
	"fmt"

	"regexp"
	"strings"

	"time"
	"unicode"

	"github.com/AlecAivazis/survey/v2"
)

// password policy
var (
	Policy = PasswordPolicyCheck{
		MinLength:           6,
		MaxLength:           64,
		RequireUpper:        true,
		RequireLower:        true,
		RequireSymbol:       true,
		RequireNumber:       true,
		DenyWhiteSpace:      true,
		AllowedSpecialChars: "{}()!@#$%^&*+_<>?:';,.][-|\"\\/",
	}
)

// operation asker function
func PasswordAskInput(prompt string, fileName string) (string, error) {

	switch strings.ToUpper(prompt) {
	case "ENCRYPT", "E":

		//ERRORTYPEerrorType := "ENCRYPTPASSERR: "

		//loop until user enter write format password for encrypt
		for {
			// fmt.Printf("Enter a password for the %s: ", prompt)
			// input, err := reader.ReadString('\n')
			var password1 string
			var password2 string

			option1 := &survey.Password{
				Message: "Enter password for Encrypt:",
			}
			survey.AskOne(option1, &password1, survey.WithIcons(func(icons *survey.IconSet) {
				icons.Question.Text = ">>"
				icons.Question.Format = "green"
			}))

			// fmt.Printf("Retype the password for the %s: ", prompt)
			//input1, err := reader.ReadString('\n')
			option2 := &survey.Password{
				Message: "Retype password for Encrypt:",
			}
			survey.AskOne(option2, &password2, survey.WithIcons(func(icons *survey.IconSet) {
				icons.Question.Text = ">>"
				icons.Question.Format = "green"
			}))

			if password1 != password2 {
				return "",
					fmt.Errorf("Password aren't matching!")
			}
			//password validation
			if err := PasswordChecker(password1, Policy); err != nil {
				fmt.Println("\nInvalid password:", err)
				fmt.Printf("Please try again!!\n")
				continue
			}
			return password1,
				nil
		}

	case "DECRYPT", "D":
		//ERRORTYPE
		errorType := "DECRYPTPASSERR: "

		//setting maxattempts for password entering
		maxAttempts := 3

		//open the file and read the data and assign data to variable
		ciphertext, saltFromFile, derivedNonce, err := FileRead(fileName, prompt)
		MainErr(errorType, err)

		//loop until user enter write format password for encrypt
		for attempts := 0; attempts < maxAttempts; attempts++ {

			var password string

			option := &survey.Password{
				Message: "Enter password for Decrypt:",
			}

			survey.AskOne(option, &password, survey.WithIcons(func(icons *survey.IconSet) {
				icons.Question.Text = ">>"
				icons.Question.Format = "green"
			}))

			//encrypting the data
			result, err := Decrypt(password, saltFromFile, derivedNonce, ciphertext)

			//success sitiuation
			if err == nil {
				//write the ciphertext data to file
				successMsg, err := FileWrite([]byte(result), fileName, nil, nil, prompt)
				return successMsg,
					err
			}
			fmt.Println("%v", err)

			//wrong password
			fmt.Println("\n****** DECRPERR: Invalid credentials or corrupt data (attempt", attempts+1, "of 3) ******")
			time.Sleep(time.Second * time.Duration(attempts+1))
		}
		//all attempts are being used so exiting the program
		return "", errors.New("Maximum attempts reached")
	}
	return "", nil
}

// password valid checker for encrypt
func PasswordChecker(password string, passwordPolicy PasswordPolicyCheck) error {
	//check password max maximum lenght
	if len(password) > passwordPolicy.MaxLength {
		return fmt.Errorf("Passoword can't be at more than %d chracters long!", passwordPolicy.MaxLength)
	}

	//checking password min minimum length
	if len(password) < passwordPolicy.MinLength {
		return fmt.Errorf("Password has to be at lease %d chracter long!", passwordPolicy.MinLength)
	}

	//checking for white space in password
	if passwordPolicy.DenyWhiteSpace && strings.ContainsAny(password, "\t\n\r") {
		return fmt.Errorf("Password can not cantain any whitespace!")
	}

	//checking for specialCharacter in password & assign a varaible for specialCharacter
	var specialCharacter string
	if passwordPolicy.RequireSymbol {
		if passwordPolicy.AllowedSpecialChars != "" {
			specialCharacter = regexp.QuoteMeta(passwordPolicy.AllowedSpecialChars)
		} else {
			specialCharacter = "{}()!@#$%^&*+_<>?:'\"\\/"
		}
	}

	//checking for chracters, numbers, symbols in password
	var HasUpper, HasLower, HasNumber, HasSymbol bool
	for _, chracter := range password {
		switch {
		case unicode.IsUpper(chracter):
			HasUpper = true
		case unicode.IsLower(chracter):
			HasLower = true
		case unicode.IsNumber(chracter):
			HasNumber = true
		case passwordPolicy.RequireSymbol && strings.ContainsRune(specialCharacter, chracter):
			HasSymbol = true
		}
	}

	var lowerRequirement []string
	if passwordPolicy.RequireUpper && !HasUpper {
		lowerRequirement = append(lowerRequirement, "Upper words")
	}
	if passwordPolicy.RequireLower && !HasLower {
		lowerRequirement = append(lowerRequirement, "Lower words")
	}
	if passwordPolicy.RequireNumber && !HasNumber {
		lowerRequirement = append(lowerRequirement, "Numbers")
	}
	if passwordPolicy.RequireSymbol && !HasSymbol {
		lowerRequirement = append(lowerRequirement, fmt.Sprintf("Password require one of these chracters.----> (%%{}^\\\"'()`~!@#$&*+_<>/-,.;')"))
	}

	if len(lowerRequirement) > 0 {
		return fmt.Errorf("Password has to be strong & contain any of these to encrypt ---->%s",
			strings.Join(lowerRequirement, ",\n"))
	}

	//checking for if password has sequential characters
	if seqIntegerCheck(password, 4) {
		return fmt.Errorf("Password contains sequential intergers. like -->(\"1111\",\"5555\",\"1234\" )")
	}
	if seqChracterCheck(password, 4) {
		return fmt.Errorf("Password contains sequential characters. like --->(\" aaa,\", \"bbb\")")
	}
	return nil
}

// password sequential integer check function
func seqIntegerCheck(word string, length int) bool {
	//if length of sequntial password integer string is shorter than provided len password is ok
	if len(word) < length {
		return false
	}
	for i := 0; i <= len(word)-length; i++ {
		isSeq := true
		for j := 1; j < length; j++ {
			if word[i+j] != word[i+j-1]+1 { //if previous byte not same to current one
				isSeq = false //reture does not has a sequential intergers and break the loop
				break
			}
		}
		if isSeq {
			return true
		}
	}
	return false
}

// password sequential check for characters
func seqChracterCheck(word string, length int) bool {
	//if length of sequntial password string is shorter than provided len password is ok
	if len(word) < length {
		return false
	}
	var currentChracter rune
	count := 1
	for _, character := range word {
		if character == currentChracter {
			count += 1
			if count >= length { /*if previous character does not equal to current one it is add
				one to count and count get equal to length if there is no sequential characters.*/
				return true
			}
		} else {
			currentChracter = character
			count = 1
		}
	}
	return false
}
