package main

import (
	"fmt"
)

func main() {
	l := largest(10, 6, 11, 45, 3)
	fmt.Println(l)
}

func largest(l ...int) int {
	var largestnum int
	for _, v := range l {
		if l > v {
		}
	}

}
