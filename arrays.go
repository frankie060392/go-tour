package main

import "fmt"

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	primesSliced := primes[1:4]
	fmt.Println(primesSliced)

	primesSliced[1] = 9
	fmt.Println(primes)

	primesSlicedDefault := primes[:4]
	fmt.Println(primesSlicedDefault)
	primesSlicedDefault = primes[1:]
	fmt.Println(primesSlicedDefault)
}
