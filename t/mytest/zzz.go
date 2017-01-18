package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	//var x uint8 = 11<<1 | 1<<5
	//fmt.Printf("%08b\n", x)
	//
	//fmt.Printf("%08b,%08b", 11<<1, 1<<5)

	var buf bytes.Buffer
	logger := log.New(&buf, "logger:", log.Ltime)
	logger.Print("Hello,log file!")

	fmt.Print(&buf)

}
