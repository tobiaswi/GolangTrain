package main

import (
	"fmt"
)

var x int

func inc() int {
	x++
	return x
}

func main() {
	fmt.Println(inc())
	fmt.Println(inc())
}
