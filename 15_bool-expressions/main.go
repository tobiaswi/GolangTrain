package main

import (
	"fmt"
)

func main() {
	x := 22
	if x > 100 {
		fmt.Println("The value is larger than 100")
	}

	if x != 0 {
		fmt.Println("The value is not 0")
	}

	if x > 0 || x < 100 {
		fmt.Println("The value is larger than 0 but smaller than 100")
	}

	if x < 50 && x > 0 {
		fmt.Println("The value between 0 and 50")
	}
}
