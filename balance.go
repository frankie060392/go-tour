package main

import (
	"fmt"
	"sync"
)

var (
	balance1 = 10
	mutex2   sync.Mutex
)

func deposit1(wg *sync.WaitGroup, amount int) {
	mutex2.Lock()
	balance1 += amount
	mutex2.Unlock()
	wg.Done()
}

func withdraw1(wg *sync.WaitGroup, amount int) {
	mutex2.Lock()
	balance1 -= amount
	mutex2.Unlock()
	wg.Done()
}

func TestMutex() {
	var wg sync.WaitGroup
	wg.Add(3)
	go deposit1(&wg, 100)
	go withdraw1(&wg, 90)
	go deposit1(&wg, 1)
	wg.Wait()
	fmt.Println(balance1)
}
