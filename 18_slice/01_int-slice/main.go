package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	greeting := []string{
		"Good morning!",
		"Bonjur!",
		"dias!",
		"Bongiorno!",
		"Selamat pagi!",
		"Guten Morgen!",
	}

	fmt.Printf("[1:2] %v\n", greeting[1:2])
	fmt.Printf("[:2] %v\n", greeting[:2])
	fmt.Printf("[5:] %v\n", greeting[5:])
	fmt.Printf("[5] %v\n", greeting[5])
	fmt.Printf("[:] %v\n", greeting[:])
}
