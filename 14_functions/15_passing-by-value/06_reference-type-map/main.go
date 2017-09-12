package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Tobi"])
}

func changeMe(z map[string]int) {
	z["Tobi"] = 32
}
