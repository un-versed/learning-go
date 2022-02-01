package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello world!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Text received %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
