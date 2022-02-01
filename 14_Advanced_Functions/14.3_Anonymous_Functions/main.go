package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World!")
	}()

	// Storing anon function return in variable
	anonFunction := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Animal Crossing")

	fmt.Println(anonFunction)
}
