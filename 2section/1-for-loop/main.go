package main

import "fmt"

func main() {
	// while style
	k := 3
	for k > 0 {
		fmt.Printf("%d ", k)
		k--
	}

	fmt.Println("--------- infinite loop ---------")

	counter := 0
	for {
		fmt.Printf("counter = %d\n", counter)
		counter++

		if counter > 5 {
			break
		}
	}

	fmt.Println("--------- skipping ---------")

	// for C style loop
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("--------- array loop ---------")

	// for array
	items := []string{"Go", "Python", "Java"}
	for k, v := range items {
		fmt.Println(k, v)
	}
}
