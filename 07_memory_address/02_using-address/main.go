package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters: ")
	fmt.Scan(&meters)
	fmt.Printf("Meteres entered: %v\nMeteres address location: %v\n", meters, &meters)
	yards := meters * metersToYards
	fmt.Println(meters,"meters is",yards,"yards.")
}