package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good Morning.",
		"Jenny": "Bonjour.",
	}

	myGreeting["Harleen"] = "Howdy."
	fmt.Println(myGreeting)
	myGreeting["Harleen"] = "Gidday."
	fmt.Println(myGreeting)

	delete(myGreeting, "Harleen")

	if val, exists := myGreeting["Harleen"]; exists {
		fmt.Println("That value exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)
}
