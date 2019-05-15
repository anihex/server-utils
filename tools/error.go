package tools

import (
	"errors"
)

// ErrIfTrue takes a boolean and checks if it's true. It returns an error
// if the value is indeed true.
func ErrIfTrue(b bool, msg string) error {
	if b {
		return errors.New(msg)
	}

	return nil
}
