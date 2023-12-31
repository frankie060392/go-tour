package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func (s *MyError) String() string {
	return fmt.Sprintf("at %v, %s", s.When, s.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

func runOk() string {
	return "It works"
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error(f float64) string {
	return fmt.Sprintf("Can not sqrt Nagative sqrt number: %v", f)
}

func Sqrt(x float64) (float64, string) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x).Error(x)
	}
	return math.Sqrt(x), "Is Ok"
}

func errorFunc() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	test := &MyError{
		time.Now(),
		"It works",
	}
	if success := runOk(); success == "It works" {
		fmt.Println(test)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
