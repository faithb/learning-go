package main

import "fmt"

func main() {
	var numbers [5]int

	numbers[0] = 10
	numbers[1] = 20

	numbers[2] = 30

	fmt.Printf("Number type %+v\n", numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Number %+v \n", numbers[i])
	}

	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3

	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	fmt.Printf("Matrix type: %+v \n", matrix)
}
