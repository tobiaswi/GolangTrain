package main

import (
	"fmt"
)

func main() {
	fmt.Println(large(10, 6, 11, 45, 3))
}

func large(l ...int) int {
	largest := l[0]
	for _, l := range l {
		if l > largest {
			largest = l
		}
	}
	return largest
}
