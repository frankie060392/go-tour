package main

import (
	"fmt"
	"math"
)

type VertexIndirection struct {
	X, Y float64
}

func (v VertexIndirection) Abs() float64 {
	v.X = 1
	v.Y = 2
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VertexIndirection) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *VertexIndirection, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func indirection() {
	fmt.Println("Pointer indirection")
	v := VertexIndirection{3, 4}
	v.Scale(2)
	fmt.Println(v)
	ScaleFunc(&v, 10)
	fmt.Println(v)
	p := &VertexIndirection{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)
	test := &VertexIndirection{4, 3}
	test.Abs()
	fmt.Println(test)

}
