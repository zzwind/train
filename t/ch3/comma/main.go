package main

import "bytes"
import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer

	i := 0
	for _, v := range s {

		buf.WriteRune(v)

		if i == 2 {
			buf.WriteByte(',')
			i = 0
		}
		i++
	}

	return buf.String()
}

func main() {
	// comma2("zzwind")
	fmt.Println(comma("12345678"))
	fmt.Println(comma2("12345678"))
}
