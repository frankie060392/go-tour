package main

import "fmt"

func def() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
