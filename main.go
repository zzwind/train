package main

import (
	"fmt"
)

func main() {
	p := 0

	pp := &p

	*pp = 1

	fmt.Println(p)
}
