package main

import "fmt"

func main() {
	// Arithmetics operators
	var sum int16 = 1 + 1
	var subtraction int16 = 1 - 1
	var division int16 = 4 / 2
	var modulo int16 = 4 % 2

	fmt.Println(sum, subtraction, division, modulo)

	var number1 int16 = 10
	var number2 int32 = 25
	var number3 int16 = 24
	var number4 int32 = 22

	// var sum2 = number1 + number2 // ❌ You can't do this
	// Go only does operations with same bitsize numbers

	var sum2 = number1 + number3
	var sum3 = number2 + number4

	fmt.Println(sum2, sum3)

	// Assignment operators
	var var1 = "String 1"
	var var2 = "String 2"

	fmt.Println(var1, var2)

	// Relational operators
	fmt.Println(1 > 2)  // false
	fmt.Println(1 < 2)  // true
	fmt.Println(1 <= 2) // true
	fmt.Println(1 >= 2) // false
	fmt.Println(1 == 2) // false
	fmt.Println(1 != 2) // true

	// Logical operators
	fmt.Println(false && true) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false

	// Unary operators
	var number int16 = 10
	number++     // 11
	number += 10 // 21

	fmt.Println(number)

	number--     // 20
	number -= 10 //10

	fmt.Println(number)

	number *= 3 // 30
	number /= 3 // 10
	number %= 2 // 0

	fmt.Println(number)

	// Ternary operators
	// ❌ There's no ternary operators in Go
}
