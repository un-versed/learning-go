package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [5]string // Every value in array should have the same type
	arr1[0] = "Position 0"
	fmt.Println(arr1)

	// Implicit declaration
	arr2 := [2]string{"Position 0", "Position 1"}
	fmt.Println(arr2)

	// Set array value = declaration items number
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr3)

	// ! Slices are not arrays, they're just a reference
	var slice = []int{1, 3, 4, 5, 6, 7, 7, 5, 4, 2}
	fmt.Println(slice) //[1 3 4 5 6 7 7 5 4 2]

	slice = append(slice, 10)
	fmt.Println(slice) // 1 3 4 5 6 7 7 5 4 2 10]

	// Copy a slice (or interval) from another array
	slice2 := arr3[1:3]
	fmt.Println(slice2) // [2 3]

	// Slices reference works like a pointer
	slice3 := arr2[0:2]
	arr2[1] = "Voldemort"
	fmt.Println(slice3, arr2) // [Position 0 Voldemort] [Position 0 Voldemort]

	// Check types
	fmt.Println(reflect.TypeOf(arr3))  //[10]int
	fmt.Println(reflect.TypeOf(slice)) // []int

	fmt.Println("---")

	// Internal arrays
	slice4 := make([]float32, 10, 11)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println("---")
	slice4 = append(slice4, 11)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	// When Go identifies that the array will reach the limit, the size is doubled
	fmt.Println("---")
	slice4 = append(slice4, 11)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
