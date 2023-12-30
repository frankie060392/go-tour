package main

import "fmt"

func pointer() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(*p)
	fmt.Println(p)
	*p = 21 // set i through the pointer
	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(j)
	p = &j // point to j
	*p = *p / 37
	fmt.Println(p)
	fmt.Println(j)

}
