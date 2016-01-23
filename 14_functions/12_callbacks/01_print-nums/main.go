package main

import "fmt"

func visit(numbers []int, postage func(int)) {
	for _, n := range numbers {
		postage(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// postage (random word) is an example of a callback
