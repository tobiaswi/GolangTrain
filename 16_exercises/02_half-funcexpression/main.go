package main

import (
	"fmt"
)

func main() {
	half := func(i int) {
		if i%2 == 0 {
			fmt.Println(i/2, true)
		} else if i%2 != 0 {
			fmt.Println(i/2, false)
		}
	}

	half(3)
	half(10)
}
