package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivisionByZero = errors.New("division by zero")
var ErrNumTooLarge = errors.New("number too large")

type OpError struct {
	Op      string
	Code    string
	Message string
	Time    time.Time
}

func (e OpError) Error() string {
	return e.Message
}

func NewOpError(op, code, message string, time time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    time,
	}
}

func doSomething() error {
	return NewOpError("doSomething", "100", "doSomething failed", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if a > 1000 {
		return 0, ErrNumTooLarge
	}

	return a / b, nil
}

func main() {

	value, err := divide(1001, 5)
	if err != nil {
		if errors.Is(err, ErrNumTooLarge) {
			fmt.Printf("The value of 1001 is %d\n", value)
		} else {
			fmt.Println(err)
		}
		return
	}

	fmt.Println(value)
}
