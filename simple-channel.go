package main

import "fmt"

func SimpleChannel() {
	// Create an unbuffered channel of type int
	c := make(chan int)

	// Start a Goroutine that sends a value on the channel
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	// Receive the value from the channel
	for i := range c {
		fmt.Println("Received value:", i)
	}

}
