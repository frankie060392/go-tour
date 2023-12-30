package main

import "fmt"

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)
	sum1 := 1
	for sum1 < 10 {
		println(sum1)
		sum1 += sum1
	}
	fmt.Println(sum1)

}
