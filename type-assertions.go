package main

import "fmt"

func typeAssertions() {
	var i interface{} = "Hello"
	var s = i.(string)
	fmt.Println(s)
	fmt.Println(i)
	s, ok := i.(string)
	fmt.Println(ok)
	f, ok := i.(int)
	fmt.Println(f, ok)
}
