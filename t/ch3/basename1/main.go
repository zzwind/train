package main

import (
	"fmt"
)

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(i)
		fmt.Println(s[i])
		if s[i] == '/' {
			fmt.Println('/')
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s

}

func main() {
	fmt.Println(basename("/a/b/c.zzz"))
}
