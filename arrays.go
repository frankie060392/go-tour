package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)
	printSlice(primes[:])

	primesSliced := primes[2:4]
	fmt.Println(primesSliced)
	printSlice(primesSliced)

	primesSliced[1] = 9
	fmt.Println(primes)

	primesSlicedDefault := primes[:4]
	fmt.Println(primesSlicedDefault)
	primesSlicedDefault = primes[1:]
	fmt.Println(primesSlicedDefault)

	var s []int
	printSlice(s)
	fmt.Println(s)
	if s == nil {
		fmt.Println("Nil!")
	}
}
