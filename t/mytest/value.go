package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	ss := [10000000]int{}

	for i := 0; i < 10000000; i++ {
		ss[i] = i
	}

	sum(ss)

	fmt.Println(time.Since(t).Seconds())

	//time.Sleep(time.Second * 2)
	tt := time.Now()

	sum2(&ss)
	zzz := &ss

	str := "zzwind"

	fmt.Println(str[0])

	strr := &str

	fmt.Println(*&strr[0])

	fmt.Println(zzz[0])
	fmt.Println(time.Since(tt).Seconds())

}

func sum(ss [10000000]int) [10000000]int {

	for i := 0; i < 10000000; i++ {
		ss[i]++
	}
	return ss
}

func sum2(ss *[10000000]int) *[10000000]int {
	for i := 0; i < 10000000; i++ {
		ss[i] += 2

	}
	return ss
}
