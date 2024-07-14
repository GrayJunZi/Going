package main

import "fmt"

func main() {
	// 循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numbers := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers: %d\n", numbers[i])
	}

	for i, number := range numbers {
		fmt.Printf("range: index %d, number %d\n", i, number)
	}

	for _, number := range numbers {
		fmt.Printf("range: ignore index, number %d\n", number)
	}

	users := map[string]int{
		"foo":   1,
		"bar":   2,
		"baz":   3,
		"alice": 4,
		"bob":   5,
	}

	for key, value := range users {
		if key == "alice" {
			break
		}
		fmt.Printf("key: %s value: %d\n", key, value)
	}
}
