package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter small number: ")
	var small int
	fmt.Scan(&small)
	fmt.Println("Enter large number: ")
	var large int
	fmt.Scan(&large)
	r := large % small
	fmt.Printf("Remainder: %v\n", r)
}
