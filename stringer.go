package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("\n%v (%v years)", p.Name, p.Age)
}

func stringer() {
	a := Person{"Frankie", 30}
	b := Person{"Test", 100}

	fmt.Println(a, b)
}
