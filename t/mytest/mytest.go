package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"time"
	"unsafe"
)

func main() {

	for i := 0; i < 10; i++ {

		st := time.Now()

		f, err := os.Open("c:/www.csdn.net.sql")

		if err != nil {
			fmt.Println(err)
		}

		buf := bufio.NewScanner(f)

		var t1, t2, t3 int

		ss1 := "# 123456 #"
		ss2 := "# 12345678 #"
		ss3 := "# 123456789 #"

		s1 := S2B(&ss1)
		s2 := S2B(&ss2)
		s3 := S2B(&ss3)

		//reg := regexp.MustCompile(`# (123456(?:78|789)?) #`)

		for buf.Scan() {

			//t := buf.Text()

			//fmt.Printf("%s", buf.Bytes())

			if bytes.Index(buf.Bytes(), s3) != -1 {
				//fmt.Println(t)
				t3++
				continue
			} else if bytes.Index(buf.Bytes(), s2) != -1 {
				t2++
				continue
			} else if bytes.Index(buf.Bytes(), s1) != -1 {
				t1++
				continue
			}

			//if strings.Index(t, "# 123456 #") != -1 {
			//	//fmt.Println(t)
			//	t1++
			//} else if strings.Index(t, "# 12345678 #") != -1 {
			//	t2++
			//} else if strings.Index(t, "# 123456789 #") != -1 {
			//	t3++
			//}

			//sub := reg.FindAllStringSubmatch(t, -1)

			//var password string
			//if sub != nil {
			//	password = sub[0][1]
			//} else {
			//	continue
			//}
			//
			//if password == "123456" {
			//	t1++
			//} else if password == "12345678" {
			//	t2++
			//} else if password == "123456789" {
			//	t3++
			//}

		}
		fmt.Println(t1, t2, t3)

		fmt.Println(time.Since(st).Seconds())
	}
}

func B2S(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

func S2B(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}
