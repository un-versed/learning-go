package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplex(write("Hello world!"), write("!Dlrow olleh"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}

func multiplex(channel1, channel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channel1:
				outputChannel <- message
			case message := <-channel2:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
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
