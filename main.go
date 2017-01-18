package main

import (
	"fmt"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	fmt.Printf("%s %[1]s %[1]s", "你好")

	s := "hello,world"

	ss := "你好世界"

	fmt.Println(s[0])
	fmt.Println(ss[0:3])

	usage := `
	
	fdsfdsfadsfdsa
	fdfsf
	dsffdsfdfdsf卧槽
	dsafads
	fas
	fdas
	fsaf
	sadfdsaf
	`

	fmt.Println(usage)
}
