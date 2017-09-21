package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Goog morning."
	myGreeting["Jenny"] = "Bonjour."
	myGreeting["Harleen"] = "Gidday."

	fmt.Println(myGreeting)
	delete(myGreeting, "Tim")
	fmt.Println(myGreeting)
}
