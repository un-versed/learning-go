package main

import (
	"errors"
	"fmt"
)

func main() {
	// Alias for types
	var type1 rune = 123456 // Alias for int32
	var type2 byte = 123    // Alias for uint8
	fmt.Println(type1)      // 123456
	fmt.Println(type2)      // 123

	// Infering types from single chars, makes Go get the value from ASCI table
	type3 := 'B'
	fmt.Println(type3) // 66

	// ZERO VALUES
	// In Go, every value has a 'default state', eg. int always start as 0
	var type4 string
	var type5 int16
	var type6 bool
	fmt.Println(type4) // ""
	fmt.Println(type5) //0
	fmt.Println(type6) // false

	// Go has a error type. The zero value for error is <nil>
	var error1 error
	var error2 error = errors.New("Error 2")
	fmt.Println(error1) // <nil>
	fmt.Println(error2) // "Error 2"

}
