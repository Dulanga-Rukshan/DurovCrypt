package DurovCrypt

import (
	"fmt"
)

func (e *ErrorHandeling) Error() string {
	return fmt.Sprintf("%s\nHelp: %s", e.Message, e.HelpMsg)
}

func NewFileError(message, help string) error {
	return &ErrorHandeling{Message: message, HelpMsg: help}
}

func MainErr(message error) {
	if message != nil {
		fmt.Printf("\nERROR: %v\n", message)
	}
}
