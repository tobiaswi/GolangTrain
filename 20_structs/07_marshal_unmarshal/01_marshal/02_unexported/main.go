package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 02}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
