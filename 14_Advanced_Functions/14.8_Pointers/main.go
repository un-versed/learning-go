package main

import "fmt"

func invertSignal(number int) int {
	return number * -1
}

func invertSignalWithPointer(number *int) *int {
	*number = *number * -1

	return number
}

func main() {
	number := 20
	invertedNumber := invertSignal(number)
	fmt.Println(invertedNumber) // -20
	fmt.Println(number)         // 20

	newNumber := 40
	fmt.Println(newNumber) // 40

	newNumber2 := invertSignalWithPointer(&newNumber)
	fmt.Println(newNumber)   // -40
	fmt.Println(*newNumber2) // -40

	*newNumber2 = 1
	fmt.Println(newNumber)   // -40
	fmt.Println(*newNumber2) // 1

}
