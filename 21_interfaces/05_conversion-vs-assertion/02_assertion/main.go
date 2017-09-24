package main

import "fmt"

func main() {
	var n interface{} = 3
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	y := 3
	fmt.Println(y + n.(int))
}
