package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Rog"]) // 46
}

func changeMe(z map[string]int) {
	z["Rog"] = 46

}
