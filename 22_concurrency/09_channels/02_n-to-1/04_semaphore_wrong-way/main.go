package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			done <- true
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// we block here until done <- true
	<-done
	<-done
	close(c)

	// to unlock we need to put the semaphore channel into its own go routine
	for n := range c {
		fmt.Println(n)
	}
}
