package main

import "fmt"

func main() {

	myFriendsName := "Bob"

	switch {
	case len(myFriendsName) == 2:
		fmt.Printf("Your name is very short %s", myFriendsName)
	case myFriendsName == "Mike":
		fmt.Println("Whats up Mike")
	case myFriendsName == "Max" && len(myFriendsName) == 3:
		fmt.Println("Hey Max, you have more than two characters in your name")
	case myFriendsName == "Bob", myFriendsName == "Sven":
		fmt.Println("Your name is either Bob or Sven")
	}
}
