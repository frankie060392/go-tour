package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	var m map[string]int
	m = make(map[string]int)
	arr := strings.Fields(s)
	fmt.Println(arr)
	for k, v := range arr {
		m[v] = len(arr[k])
	}
	return m
}

func testWorldCount() {
	m := wordCount("Frankie tran")
	fmt.Println(m)
}
