package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func replace(str string) string {
	return ""
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	zz := "wind"
	fmt.Println(replace(zz))
	fmt.Println(zz)
	fmt.Println(time.Now())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println(0 - r.Intn(12))

	dict := make(map[string]int)

	dict["zz"] = 1
	dict["zzwind"] = 2

	changedict(dict)

	fmt.Println(dict)
	fmt.Println("我是帅哥")
	fmt.Println("hahahahaha")
	fmt.Println()

	fmt.Println(utf8.RuneCountInString("我是帅哥aa"))

	zzz := "我是帅哥"
	fmt.Println(zzz[0])

}

func changedict(d map[string]int) {
	d["name"] = 3
}
