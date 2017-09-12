package main

import (
	"fmt"
)

func main() {
	d := []float64{43, 56, 87, 12, 45, 57}
	a := average(d)
	fmt.Println(a)
}

func average(sf []float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
