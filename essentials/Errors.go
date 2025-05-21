package DurovCrypt

import (
	"fmt"
	"os"
)

func (e *ErrorHandeling) Error() string {
	return fmt.Sprintf("%s\nHelp: %s", e.Message, e.HelpMsg)
}

func NewFileError(message, help string) error {
	return &ErrorHandeling{Message: message, HelpMsg: help}
}

func MainErr(errorType string, message error) {
	if message != nil {
		fmt.Printf("%v %v\n", errorType, message)
		os.Exit(0)
	}
}
