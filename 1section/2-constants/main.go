package main

import "fmt"

const HOST string = "localhost"

func main() {
	fmt.Println(HOST)

	const AppName string = "GO"
	fmt.Printf("Hello, %s!\n", AppName)

	const pi float64 = 3.1415926
	fmt.Println(pi)
}
