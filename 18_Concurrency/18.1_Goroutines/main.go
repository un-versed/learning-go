package main

import (
	"fmt"
	"time"
)

func main() {
	go writeLoop("Hello world!")
	writeLoop("Hey Alexa!")
}

func writeLoop(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second / 2)
	}
}
