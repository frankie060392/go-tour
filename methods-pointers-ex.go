package main

import (
	"fmt"
	"math"
)

type VertexPoinerEx struct {
	X, Y float64
}

func Abs(v VertexPoinerEx) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scalez(v *VertexPoinerEx, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodPoinerEx() {
	v := VertexPoinerEx{3, 4}
	Scalez(&v, 10)
	fmt.Println(Abs(v))
}
