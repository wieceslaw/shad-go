//go:build !solution

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int64)

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		str := string(data)
		words := strings.Split(str, "\n")

		for _, word := range words {
			counts[word]++
		}
	}

	for key, value := range counts {
		if value >= 2 {
			fmt.Printf("%d\t%s\n", value, key)
		}
	}
}
