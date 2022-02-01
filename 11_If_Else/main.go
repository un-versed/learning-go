package main

import "fmt"

func main() {
	number := 10

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Lesser than 0")
	}

	// Variables created in if condition, only have the if scope
	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Greather than 0")
	} else {
		fmt.Println("Lesser than 0")
	}
}
