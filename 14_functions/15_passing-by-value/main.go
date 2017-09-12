package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) // Memory Address

	changeMe(&age)
	fmt.Println(&age) // Memory Address
	fmt.Println(age)  // Value at Address
}

func changeMe(z *int) {
	fmt.Println(z)  // Memory Address
	fmt.Println(*z) // Value at Address
	*z = 24
	fmt.Println(z)  //Memory Address
	fmt.Println(*z) // Value at Address
}
