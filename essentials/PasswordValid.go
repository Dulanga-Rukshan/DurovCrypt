package DurovCrypt

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// password valid checker for encrypt
func PasswordChecker(password string, passwordPolicy PasswordPolicyCheck) error {
	//check password max maximum lenght
	if len(password) > passwordPolicy.MaxLength {
		return fmt.Errorf("Passoword can't be at more than %d chracters long!", passwordPolicy.MaxLength)
	}

	//checking password min minnimum length
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
		return fmt.Errorf("Password has to be strong & contain any of these to encrypt ---->%s", strings.Join(lowerRequirement, ",\n"))
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
	if len(word) < length {
		return false
	}
	var currentChracter rune
	count := 1
	for _, character := range word {
		if character == currentChracter {
			count += 1
			if count >= length { // if previous chracter does not equal to current one it is add one to count and count get equal to length if there is no sequential characters.
				return true
			}
		} else {
			currentChracter = character
			count = 1
		}
	}
	return false
}
