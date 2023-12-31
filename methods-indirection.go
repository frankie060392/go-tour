package main

import "fmt"

type VertexIndirection struct {
	X, Y float64
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
}
