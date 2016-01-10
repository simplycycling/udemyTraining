package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("type\tbinary\t\t\t\t\t\tdecimal\n")
	fmt.Printf("KB\t")
	fmt.Printf("%b\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("MB\t")
	fmt.Printf("%b\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("GB\t")
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("TB\t")
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)

}
