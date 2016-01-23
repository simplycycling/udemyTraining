package main

import "fmt"

func filter(numbers []int, example func(int) bool) []int {
	xs := []int{} // the idiomatically correct way to say this would be "var xs []int"
	for _, n := range numbers {
		if example(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T", example)
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}
