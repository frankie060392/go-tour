package main

import (
	"fmt"
	"math"
)

type VertexMethod struct {
	X, Y float64
}

func (v VertexMethod) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v VertexMethod) test() int {
	return int(v.X) + int(v.Y)
}

func methods() {
	v := VertexMethod{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.test())
}
