package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

var c = make(chan string, 10)

func main() {

	st := time.Now()

	var t1, t2, t3 int

	//var lines []string

	go readfile()

	reg := regexp.MustCompile(`# (123456(?:78|789)?) #`)
	var i = 0
	for _ = range c {
		i++
		fmt.Println(i)
		sub := reg.FindAllStringSubmatch(<-c, -1)

		var password string
		if sub != nil {
			password = sub[0][1]
		} else {
			continue
		}

		//_ = strings.Split(t, " # ")[1]
		//
		if password == "123456" {
			t1++
		} else if password == "12345678" {
			t2++
		} else if password == "123456789" {
			t3++
		}

	}
	fmt.Println(t1, t2, t3)
	fmt.Scanln()
	fmt.Println(time.Since(st).Seconds())
}

func readfile() {
	f, err := os.Open("c:/www.csdn.net.sql")

	if err != nil {
		fmt.Println(err)
	}

	buf := bufio.NewScanner(f)

	for buf.Scan() {
		c <- buf.Text()
	}
}

func regtext() {

}
