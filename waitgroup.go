package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex1  sync.Mutex
)

func init() {
	balance = 100
}

func deposit(amount int, wg *sync.WaitGroup) {
	mutex1.Lock()
	balance += amount
	mutex1.Unlock()
	wg.Done()
}

func withdraw(amount int, wg *sync.WaitGroup) {
	mutex1.Lock()
	balance -= amount
	mutex1.Unlock()
	wg.Done()
}

func WaitGroupTest() {
	var wg sync.WaitGroup
	wg.Add(3)
	go deposit(20, &wg)
	go withdraw(40, &wg)
	go deposit(100, &wg)
	wg.Wait()
	fmt.Println(balance)
}
