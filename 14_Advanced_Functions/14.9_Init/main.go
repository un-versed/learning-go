package main

import "fmt"

// Init function is executed before main()

func main() {
	fmt.Println("Main function being executed")
}

func init() {
	fmt.Println("Init function beign executed")
}
