package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Sunday"
	fmt.Println("Today is ", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend! No work!")
	case "Monday", "Tuesday":
		fmt.Println("Works day. Lots of meeting")
	default:
		fmt.Println("Mid-week")
	}

	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", v)
		}
	}

	checkType(2)
	checkType("hello")
	checkType(3.21)
}
