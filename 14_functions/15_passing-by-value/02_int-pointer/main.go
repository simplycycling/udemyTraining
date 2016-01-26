package main

import "fmt"

func main() {

	age := 46         // assigns 46 to age
	fmt.Println(&age) // prints memory address of age

	changeMe(&age) // runs changeMe function

	fmt.Println(&age) // prints memory address of age
	fmt.Println(age)  // prints age
}

func changeMe(z *int) { // func changeMe takes input z (&age), type pointer to mem address of int
	fmt.Println(z) // prints z, which is the mem address of
	fmt.Println(*z)
	*z = 24
	fmt.Println(z)
	fmt.Println(*z)
}
