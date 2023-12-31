package main

import (
	"fmt"
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

func errorFunc() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
