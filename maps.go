package main

import "fmt"

type VertexMap struct {
	Lat, Long float64
}

var m map[string]VertexMap

var n = map[string]VertexMap{
	"Bell Labs": {
		1, 3,
	},
}

func maps() {
	fmt.Println(m == nil)
	m = make(map[string]VertexMap)
	m["Bell Labs"] = VertexMap{
		1.1, -2,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(n["Bell Labs"])
	delete(n, "Bell Labs")
	fmt.Println(n["Bell Labs"])

	elem, ok := m["Bell Labs"]
	fmt.Println("%b and %b", elem, ok)
}
