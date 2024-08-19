package utils

import "fmt"

func AppendError(exitErr, newErr error) error {
	if exitErr == nil {
		return newErr
	}
	return fmt.Errorf("%v, %w", exitErr, newErr)
}
