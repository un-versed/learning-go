package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Program recovered from panic with success!")
	}
}

func approvedStudent(n1, n2 float32) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	// When panic is called, the program stops and run all defers
	panic("The average is exactly 6")
}

func main() {
	fmt.Println(approvedStudent(6, 6))
	fmt.Println("Post execution")
}
