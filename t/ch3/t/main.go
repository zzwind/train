package main

import "fmt"

func main() {
	s := "hello, 世界"

	fmt.Println(s[11])
	// fmt.Println(len(s))
	// fmt.Println(utf8.RuneCountInString(s))

	// for i := 0; i < len(s); {
	// 	fmt.Println(s[i:])
	// 	r, size := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%d\t%c\n", i, r)
	// 	i += size
	// }
}
