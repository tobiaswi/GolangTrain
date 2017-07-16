package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// noacess to x
	//this does not compile
	fmt.Println(x)
}
