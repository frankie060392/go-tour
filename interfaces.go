package main

import (
	"fmt"
	"math"
)

type VertexInterface struct {
	X, Y float64
}

type MyFloatInterface float64

type Abser interface {
	Abs() float64
}

func (f MyFloatInterface) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *VertexInterface) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func interfaces() {
	var a Abser
	f := MyFloatInterface(-math.Sqrt2)
	v := VertexInterface{3, 4}
	a = f  // a MyFloat implements Abser
	a = &v // a *Vertexinterface implements Abser
	// a = v // Error
	fmt.Println(a.Abs())
}
