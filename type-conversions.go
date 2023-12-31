package main

import (
	"fmt"
	"math"
)

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*x))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
