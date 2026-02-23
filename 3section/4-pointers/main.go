package main

import "fmt"

func main() {
	num := 10
	modifyValue(&num)

	fmt.Println(num)
}

func modifyValue(val *int) {
	*val = *val * 10
	fmt.Printf("val: %d\n", *val)
}
