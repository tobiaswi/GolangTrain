package main

import (
	"fmt"
)

//Contact is used to as assertion example
type Contact struct {
	greeting string
	name     string
}

//SwitchOnType this is an assert; asserting, "x is of this type"
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {

}
