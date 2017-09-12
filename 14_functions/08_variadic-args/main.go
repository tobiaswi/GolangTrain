package main

import (
	"fmt"
)

func main() {
	d := []float64{43, 56, 87, 12, 45, 57}
	n := average(d...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
