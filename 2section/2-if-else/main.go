package main

import "fmt"

func main() {
	score := 85

	if score > 85 {
		fmt.Println("Grade: A")
	} else if score > 70 {
		fmt.Println("Grade: B")
	} else if score > 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	userAccess := map[string]bool{
		"hung":  false,
		"faith": true,
	}

	if hasAccess, ok := userAccess["hung"]; ok && hasAccess {
		fmt.Println("Hung can access system")
	} else {
		fmt.Println("access not granted")
	}
}
