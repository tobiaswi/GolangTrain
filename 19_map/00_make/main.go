package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjur."

	fmt.Println(myGreeting)
}
