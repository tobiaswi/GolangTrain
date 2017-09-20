package main

import "fmt"

func main() {
	z := (true && false) || (false && true) || !(false && false)
	fmt.Printf("%v", z)
}
