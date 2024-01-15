package main

import (
	"fmt"
	"math/rand"
	"runtime"
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

func sender(c chan<- int, name string) {
	for i := 1; i <= 100; i++ {
		c <- 1
		fmt.Printf("%s has sent 1 to channel\n", name)
		runtime.Gosched()
	}
}

func publisher() <-chan int {
	c := make(chan int)

	go func() {
		for i := 1; i <= 1000; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func consumer(c <-chan int, name string) {
	counter := 0

	for value := range c {
		fmt.Printf("Consumer %s is doing task %d\n", name, value)
		counter++
		time.Sleep(time.Millisecond * 20)
	}

	fmt.Printf("Consumer %s has finished %d task(s)\n", name, counter)
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
	myChan1 := make(chan int)

	go sender(myChan1, "S1")
	go sender(myChan1, "S2")
	go sender(myChan1, "S3")

	start := 0

	for {
		start += <-myChan1
		fmt.Println(start)

		if start >= 300 {
			break
		}
	}
	myChan2 := publisher()
	maxConsumer := 5

	for i := 1; i <= maxConsumer; i++ {
		go consumer(myChan2, fmt.Sprintf("%d", i))
	}

	time.Sleep(time.Second * 10)
}
