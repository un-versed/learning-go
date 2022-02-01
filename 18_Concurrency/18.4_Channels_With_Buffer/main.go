package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)

	channel <- "Hello world!"
	channel <- "Hello world 2!"

	text := <-channel
	text2 := <-channel

	fmt.Println(text)
	fmt.Println(text2)
}
