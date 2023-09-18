package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	trimInput := strings.Trim(input, "\r\n")
	var start, end int
	for next := 1; next < len(trimInput); next++ {
		hasFound := true
		for i := start; i < (end + 1); i++ {
			if trimInput[i] == trimInput[next] {
				start = i + 1
				hasFound = false
				break
			}
		}

		end = next

		if hasFound && end-start+1 == 14 {
			fmt.Println(end + 1)
			break
		}
	}
}
