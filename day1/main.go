package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	trim_input := strings.Trim(input, "\r\n")

	max := 0
	elfs := strings.Split(trim_input, "\r\n\r\n")
	for _, elf := range elfs {
		foods := strings.Split(elf, "\r\n")
		total := 0
		for _, cal := range foods {
			cal_int, err := strconv.Atoi(cal)
			if err != nil {
				fmt.Println(err)
				return
			}

			total += cal_int
		}

		if total > max {
			max = total
		}
	}

	fmt.Println(max)
}
