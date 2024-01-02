package main

type List[T any] struct {
	next *List[T]
	val  T
}

type Number[T any] struct {
	Value T
}
