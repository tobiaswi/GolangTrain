package main

import "fmt"

func main() {
	greet("John", "Smith")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
