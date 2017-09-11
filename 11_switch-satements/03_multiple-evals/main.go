package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Whats up Daniel")
	case "Medhi", "Julian":
		fmt.Println("Whats up Medhi, or Julian")
		fallthrough
	case "Jenny":
		fmt.Println("Whats up Jenny")
	case "Julius":
		fmt.Println("Whats up Julius")
	case "Mike":
		fmt.Println("Wahts up Mike")
		fallthrough
	default:
		fmt.Println("Dont you know someone?")
	}
}
