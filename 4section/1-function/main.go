package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(x int, y int) {
	fmt.Printf("add: %d + %d = %d\n", x, y, x+y)
}

func calculateArea(width, height float64) float64 {
	if width < 0 || height < 0 {
		fmt.Println("Error: width and height must be positive")
		return 0
	}

	return width * height
}

func main() {
	greet("Bob")
	add(1, 2)
	area := calculateArea(-10, 20)
	fmt.Println(area)
}
