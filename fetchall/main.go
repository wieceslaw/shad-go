//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := os.Args[1:]

	ch := make(chan any)
	start := time.Now()

	for _, url := range urls {
		go func() {
			defer func() { ch <- 1 }()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%v\t%d\t%s", time.Since(start), len(body), url)
		}()
	}

	for range urls {
		<-ch
	}

	fmt.Printf("Elapsed %v", time.Since(start))
}
