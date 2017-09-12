package main

import "fmt"

func main() {
	greet("John", "Smith")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
