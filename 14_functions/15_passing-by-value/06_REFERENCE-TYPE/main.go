package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMe(m)
	fmt.Println(m) // [Rog]
}

func changeMe(z []string) {
	z[0] = "Rog"
	fmt.Println(z) // Fuck y'all, it's [Rog}!

}
