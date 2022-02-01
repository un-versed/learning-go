package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// Common for
	for i < 10 {
		i++
		time.Sleep((time.Second / 2))
		fmt.Println("Incrementing i", i) // Incrementing i 1..2..3...
	}

	fmt.Println("i", i)

	// Inline for
	for j := 0; j < 10; j++ {
		time.Sleep(time.Millisecond)
		fmt.Println("Incrementing j", j) // Incrementing j 1..2..3...
	}

	// For loop with range
	names := [3]string{"John", "Mary", "Jesus"}

	for index, name := range names {
		fmt.Println(index, name) // John Mary Jesus
	}

	for _, letter := range "Potato" {
		fmt.Println(string(letter)) // p o t a t o
	}

	// For loop with maps
	user := map[string]string{
		"name": "Gustavo",
		"age":  "22",
	}

	for key, value := range user {
		fmt.Println(key, value) // name Gustavo \r\n age 22
	}
}
