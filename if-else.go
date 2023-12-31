package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		fmt.Println("x is unsigned int")
	}
	return fmt.Sprint(math.Sqrt(x))
}

func ifelse() {
	fmt.Println(sqrt(2), sqrt(-4))
}
