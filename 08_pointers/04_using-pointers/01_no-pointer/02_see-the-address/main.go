package main

import "fmt"

func zero(z int) {
	// address in func zero
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
	// set z to zero
	z = 0
}

func main() {

	x := 5

	// address in func main
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)

	// pass x to zero func
	zero(x)

	// x is still 5
	fmt.Println(x)
}
