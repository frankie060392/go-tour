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

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods() {
	v := VertexMethod{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.test())

	f := MyFloat(-math.Sqrt(4))
	fmt.Println(f)
	fmt.Println(f.Abs())
}
