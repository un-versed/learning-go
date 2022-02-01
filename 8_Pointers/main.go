package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2) // 10 10

	var1++
	fmt.Println(var1, var2) // 11 10

	// Pointers are a reference in the memory
	var var3 int
	var var4 *int

	fmt.Println(var3, var4) // 0 <nil>

	var3 = 120
	var4 = &var3

	fmt.Println(var3, *var4, var4)

	var3 = 150
	fmt.Println(var3, *var4, var4)
}
