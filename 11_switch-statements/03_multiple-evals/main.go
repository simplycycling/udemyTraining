package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, and or Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both your names start with M! Derp!")
	case "Julien", "Sushant":
		fmt.Println("Hi")
	}
}
