package main

import "fmt"

func main() {
	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(<-(factorial(4)))
		}
	}()
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
