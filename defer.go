package main

import "fmt"

func def() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
