package main

import "fmt"

func main() {
	if false {
		fmt.Println("first print statement")
	} else if false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("thrid print statment")
	} else {
		fmt.Println("fourth print statment")
	}
}
