package main

import (
	"fmt"

	"github.com/tobiaswi/GolangTrain/02_packages/printname"
)

func main() {
	fmt.Println(printname.PrintName("Hi Tobi!"))
	fmt.Println(printname.MyName)
}
