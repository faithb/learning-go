package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Dan":   60,
	}

	fmt.Printf("StudentGrades: %+v\n", studentGrades)

	studentGrades["Alice"] = 95
	fmt.Printf("StudentGrades: %+v\n", studentGrades)

	alice, ok := studentGrades["Dan"]

	if ok {
		fmt.Printf("Dan: %+v\n", alice)
	}

	key := "Jame"

	if value, ok := studentGrades[key]; ok {
		fmt.Printf("%s : %+v\n", key, value)
	}

	delete(studentGrades, "Alice")

	fmt.Printf("StudentGrades: %+v\n", studentGrades)

	configs := make(map[string]int)
	fmt.Printf("configs: %+v %T\n", configs, configs)

	if configs == nil {
		fmt.Printf("Configs is nil")
	}
}
