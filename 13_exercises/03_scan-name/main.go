package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter your name: ")
	var input string
	fmt.Scan(&input)
	fmt.Printf("Hello %s \n", input)
}
