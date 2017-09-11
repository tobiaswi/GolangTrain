package main

import "fmt"

func main() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			fmt.Printf("%v is divisable by three\n", i)
		} else {
			fmt.Printf("%v is is not divisable by three\n", i)
		}
	}
}
