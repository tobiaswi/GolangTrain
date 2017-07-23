package main

import (
	"fmt"
)

func main() {

	x := 15 % 3
	fmt.Println(x)

	if x != 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 30; i++ {
		if i%4 != 0 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
