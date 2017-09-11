package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Whats up Daniel")
	case "Medhi":
		fmt.Println("Whats up Medhi")
		fallthrough
	case "Jenny":
		fmt.Println("Whats up Jenny")
		fallthrough
	default:
		fmt.Println("Dont you know someone?")
	}
}
