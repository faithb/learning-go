package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("Function SimpleDefer: Start")

	defer fmt.Println("Function SimpleDefer: Deferred")
	fmt.Println("Function SimpleDefer: Middle")
	fmt.Println("Function SimpleDefer: Middle")
	fmt.Println("Function SimpleDefer: Middle")
	fmt.Println("Function SimpleDefer: Middle")
}

func lifoSimpleDefer() {
	fmt.Println("Function SimpleDefer: Start")

	defer fmt.Println("First Function SimpleDefer: Deferred")
	defer fmt.Println("Second Function SimpleDefer: Deferred")
	fmt.Println("Function SimpleDefer: Middle")
}

func main() {

	file, err := os.Create("defer.text")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	simpleDefer()

	lifoSimpleDefer()

	fmt.Println("Last of main")
}
