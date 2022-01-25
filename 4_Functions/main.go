package main

import (
	"fmt"
)

func sum(a int8, b int8) int8 {
	return a + b
}

// Functions can return multiple values
func multiOperations(n1, n2 int8) (int8, int8) {
	sum := sum(n1, n2)
	sub := n1 - n2

	return sum, sub
}

func main() {
	result := sum(4, 5)
	fmt.Println(result) // 9

	// Declaring a function with variable
	var print = func(text string) string {
		fmt.Println(text)

		return text
	}
	print("Hi folks@!") // "Hi folks@!"

	sumResult, subResult := multiOperations(10, 5)
	fmt.Println(sumResult) // 15
	fmt.Println(subResult) // 5

	// Ignore specific returns from function
	sumResult2, _ := multiOperations(11, 2)
	fmt.Println(sumResult2) // 13
}
