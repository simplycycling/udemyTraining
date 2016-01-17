package main

import "fmt"

func main() {

	myFriendsName := "Ma"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
		fallthrough
	case myFriendsName == "Ma":
		fmt.Println("Wassup Tim")
		fallthrough
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
		fallthrough
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
		fallthrough
	case myFriendsName == "Julien":
		fmt.Println("Whats up Julien")
		fallthrough
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
		fallthrough
	default:
		fmt.Println("nothing matched, this is the default")
	}
}
