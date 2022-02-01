package main

import "fmt"

// Functions can return multiple values without naming it
func multiOperations(n1, n2 int8) (sum int8, sub int8) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	sum, sub := multiOperations(3, 2)

	fmt.Println(sum) // 5
	fmt.Println(sub) // 1
}
