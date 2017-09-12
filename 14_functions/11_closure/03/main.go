package main

import "fmt"

func main() {
	var x int
	inc := func() int {
		x++
		return x
	}
	fmt.Println(inc())
	fmt.Println(inc())
}
