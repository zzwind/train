package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil
		}
		log.Printf("server not responding (%s); retyring...", err)
		time.Sleep(time.Second << uint(tries))
	}

	io.EOF

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)

}
