package main

import "fmt"

func fibo(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x = y
		y = x + y
		c <- y
	}
	close(c)
}

func channelsFibo() {
	c := make(chan int, 10)
	go fibo(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
