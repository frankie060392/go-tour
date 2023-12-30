package main

import "fmt"

const Pi = 3.14

func constant() {
	const World = "Test"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day")

	const Truth = true

	fmt.Println("Go rules ?", Truth)
}
