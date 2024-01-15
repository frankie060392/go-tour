package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test(c chan int) {
	go func() {
		for i := 1; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
}

func ChannelTestAdvance() {

	var myChan = make(chan int)
	// go func() {
	// 	for i := 1; i < 10; i++ {
	// 		myChan <- i
	// 	}
	// 	close(myChan)
	// }()
	test(myChan)
	for {
		value, isAlive := <-myChan
		if !isAlive {
			fmt.Println("Channel has been closed")
			break
		}
		fmt.Println("Value is: ", value)
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ts := time.Second * time.Duration(r.Intn(5))
		fmt.Println(ts, "ch1")
		time.Sleep(ts)
		ch1 <- 1
	}()

	go func() {
		ts := time.Second * time.Duration(r.Intn(5))
		fmt.Println(ts, "ch2")
		time.Sleep(ts)
		ch2 <- 2
	}()

	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)

	select {
	case v1 := <-ch1:
		fmt.Println("Ch1 come first with value:", v1)
		fmt.Println("then ch2 with value:", <-ch2)
	case v2 := <-ch2:
		fmt.Println("Ch2 come first with value:", v2)
		fmt.Println("then ch2 with value:", <-ch1)
	}
}
