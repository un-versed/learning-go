package main

import "fmt"

func func1() {
	fmt.Println("Function 1")
}

func func2() {
	fmt.Println("Function 2")
}

// A defer (defers) the execution of a function
// Case of use: you can use a defer to close a database connection, because the defer clausule delays the execution
func main() {
	defer func1()
	func2()
}
