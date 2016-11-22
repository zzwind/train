package main

import (
	"bytes"
	"fmt"
)

func intsToString(value []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range value {
		if i > 0 {
			fmt.Println(",")
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
