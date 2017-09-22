package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func (p person) greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) greeting() {
	fmt.Println("Miss MoneyPenny, so good to see you.")
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	p1.greeting()
	p1.person.greeting()
	p2.greeting()
	p2.person.greeting()
}
