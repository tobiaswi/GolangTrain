package main

import (
	"fmt"
)

func half(i int) (int, bool) {
	x := i
	d := (i / 2)
	if x%2 == 0 {
		return d, true
	}
	return d, false
}

func main() {
	fmt.Println(half(11))
}
