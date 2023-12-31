package main

type VertexPointer struct {
	X int
	Y int
}

func pointerVer() {
	v := VertexPointer{1, 2}
	p := &v
	println(p)
	println((*p).X)
	println(p.X)
}
