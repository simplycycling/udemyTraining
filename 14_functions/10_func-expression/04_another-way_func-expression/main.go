package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "hello world"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
