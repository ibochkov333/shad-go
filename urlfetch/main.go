//go:build !solution

package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		resp, err := http.Get(arg)
		if err != nil {
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(string(b))
	}
}
