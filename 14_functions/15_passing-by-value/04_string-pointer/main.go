package main

import "fmt"

func main() {

	name := "Rog"
	fmt.Println(&name) // mem address

	changeMe(&name)

	fmt.Println(&name) // mem address
	fmt.Println(name)  // Rocky
}

func changeMe(z *string) {
	fmt.Println(z)  // mem address
	fmt.Println(*z) // Rog
	*z = "Rocky"
	fmt.Println(z)  // mem address
	fmt.Println(*z) // Rocky

}
