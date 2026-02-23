package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie"}

	fmt.Println(names)

	items := make([]int, 3, 4)

	fmt.Printf("Items: %+v, Len: %d, Cap: %d \n", items, len(items), cap(items))

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	items = append(items, 4)
	items = append(items, 5)
	items = append(items, 5)
	items = append(items, 5)
	items = append(items, 5)

	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
}
