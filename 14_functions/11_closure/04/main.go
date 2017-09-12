package main

import "fmt"

func wrapper() func() int {
	var n int
	return func() int {
		n++
		return n
	}
}

func main() {
	w := wrapper()
	fmt.Println(w())
	fmt.Println(w())
}
