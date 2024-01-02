package main

import (
	"fmt"
	"time"
)

func fiboSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}
func selectChan() {
	fmt.Println("Select channels")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fiboSelect(c, quit)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-boom:
			fmt.Println("Boom.")
			return
		default:
			fmt.Println("Waiting.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
