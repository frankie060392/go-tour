package main

import "fmt"

type VertexMap struct {
	Lat, Long float64
}

var m map[string]VertexMap

func maps() {
	m = make(map[string]VertexMap)
	m["Bell Labs"] = VertexMap{
		1.1, -2,
	}
	fmt.Println(m["Bell Labs"])
}
