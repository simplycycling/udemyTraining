package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	total := 0.0
	// a better way to do the above line is "var total float64". It's idiomatic, and sets the value to 0.
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
