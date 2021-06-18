package errorhandling

import (
	"errors"
	"fmt"
)

func IsPositive(arg int) (int, error) {
	if arg <= 0 {
		return -1, errors.New("Its a Negative")
	}
	return arg, nil
}

type errorMessage struct {
	arg int
	msg string
}

func (e errorMessage) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.msg)
}

func IsPositive1(arg int) (int, error) {
	if arg <= 0 {
		return -1, &errorMessage{arg, "Its a Negative"}
	}
	return arg, nil
}
