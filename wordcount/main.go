//go:build !solution

package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {
	args := os.Args[1:]
	mp := make(map[string]int)

	for i := 0; i < len(args); i++ {
		f, err := os.Open(args[i])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			mp[scanner.Text()]++
		}
	}

	for key, val := range mp {
		if val >= 2 {
			fmt.Printf("%d\t%s\n", val, key)
		}
	}
}
