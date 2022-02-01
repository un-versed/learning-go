package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		writeLoop("Hello world!")
		waitGroup.Done()
	}()

	go func() {
		writeLoop("Hey Alexa!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func writeLoop(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second / 2)
	}
}
