package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	// 2 | 32
	//
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%03b\n", 2)

	fmt.Printf("\n%b\n", 10)
	fmt.Printf("%b\n", 1<<1)
	fmt.Printf("%b\n", 1<<5)
}
