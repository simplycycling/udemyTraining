package main

import "fmt"

func main() {
	age := 46
	changeMe(age)
	fmt.Println(age)
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}
