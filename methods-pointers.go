package main

import (
	"fmt"
	"math"
)

type VertexMethodPointer struct {
	X, Y float64
}

func (v VertexMethodPointer) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VertexMethodPointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodPointers() {
	v := VertexMethodPointer{3, 4}
	fmt.Println(v.Abs())
	v.Scale(2)
	fmt.Println(v.Abs())
	fmt.Println(v.X, v.Y)
}
