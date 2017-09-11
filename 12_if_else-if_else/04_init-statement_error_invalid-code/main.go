package main

import "fmt"

func main() {
	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	fmt.Println(food) //the scope of food is contained inside its if statement
}
