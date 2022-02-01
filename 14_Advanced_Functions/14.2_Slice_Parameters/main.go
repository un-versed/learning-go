package main

import "fmt"

// Functions can use slices as parameters
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(total)
}
