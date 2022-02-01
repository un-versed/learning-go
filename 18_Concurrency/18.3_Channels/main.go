package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	channel2 := make(chan string)

	go writeLoop("Hello world!", channel)
	go writeLoop("Another Hello world!", channel2)

	// Using a range to get values from channel
	for text := range channel2 {
		fmt.Println(text)
	}

	// Using an infinite for loop to get values from channel
	for {
		text, open := <-channel
		if !open {
			break
		}
		fmt.Println(text)
	}

}

func writeLoop(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second / 2)
	}

	close(channel)
}
