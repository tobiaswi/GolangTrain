package main

import (
	"fmt"
)

const (
	// A 0
	A = iota
	// B 1
	B = iota
	// C 2
	C = iota
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}