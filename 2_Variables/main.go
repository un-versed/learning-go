package main

import "fmt"

func main() {
	// Single line declaration
	var var1 string = "Var 1" // (String) Declaring type (explicit)
	var var2 int = 2          // (int) Declaring type (explicit)
	var3 := "Var 2"           // (String) without declaration (implicit)
	var4 := 2                 // (int) without declaration (implicit)

	// Multiple declaration (explicit)
	var (
		var5  string
		var6  bool
		var7  int
		var8  int16
		var9  int32
		var10 int64
		var11 float32
		var12 float64
		var13 []int
	)

	// Multiple declarations (implicit)
	var14, var15, var16 := "Multiple vars", 3, true

	// Set values
	var13 = []int{1, 2, 3}

	// Invert two variables
	var (
		varA string = "I'm A"
		varB string = "I'm B"
	)

	varA, varB = varB, varA

	// Print
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)
	fmt.Println(var7)
	fmt.Println(var8)
	fmt.Println(var9)
	fmt.Println(var10)
	fmt.Println(var11)
	fmt.Println(var12)
	fmt.Println(var13)
	fmt.Println(var14)
	fmt.Println(var15)
	fmt.Println(var16)
}
