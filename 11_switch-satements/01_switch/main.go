package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Whats up Daniel")
	case "Medhi":
		fmt.Println("Whats up Medhi")
	case "Jenny":
		fmt.Println("Whats up Jenny")
	default:
		fmt.Println("Dont you know someone?")
	}
}
