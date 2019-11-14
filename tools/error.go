package tools

import "fmt"

// ErrIfTrue takes a boolean and checks if it's true. It returns an error
// if the value is indeed true.
func ErrIfTrue(b bool, msg string, params ...interface{}) error {
	if b {
		return fmt.Errorf(msg, params...)
	}

	return nil
}
