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

const (
	// D 0
	D = iota
	// E 1
	E = iota
	// F 2
	F = iota
)

const (
	_ = iota
	// G 1
	G = iota
	// H 2
	H = iota
)

const (
	// throwing away 0
	_ = iota
	// KB with iota
	KB = 1 << (iota * 10)
	// MB with iota
	MB = 1 << (iota * 10)
	// GB with iota
	GB = 1 << (iota * 10)
	// TB with iota
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

	fmt.Println(G)
	fmt.Println(H)

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t%d\n", KB, KB)
	fmt.Printf("%b\t%d\n", MB, MB)
	fmt.Printf("%b\t%d\n", GB, GB)
	fmt.Printf("%b\t%d\n", TB, TB)
}
