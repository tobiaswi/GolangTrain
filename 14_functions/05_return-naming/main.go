package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("John", "Smith"))
}

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
