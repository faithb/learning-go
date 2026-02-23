package main

import (
	"fmt"
	"strings"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]
	return
}

func main() {
	value, err := divide(10, 1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	firstName, lastName := splitName("Joshe Adam")

	fmt.Printf("FirstName: %s \nLastName: %s \n", firstName, lastName)
}
