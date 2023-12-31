package main

import "fmt"

func fibonacci() func(int) int {
	result := 0
	first := 0
	second := 1
	return func(i int) int {
		if i == 0 || i == 1 {
			return i
		}
		result = first + second
		first = second
		second = result
		return result
	}
}

func closureFibo() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
