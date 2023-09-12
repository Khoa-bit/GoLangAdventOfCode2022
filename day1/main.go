package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string) int {
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
				return -1
			}

			total += cal_int
		}

		if total > max {
			max = total
		}
	}

	return max
}

func part2(input string) int {
	trim_input := strings.Trim(input, "\r\n")

	max := make([]int, 3)
	elfs := strings.Split(trim_input, "\r\n\r\n")
	for _, elf := range elfs {
		foods := strings.Split(elf, "\r\n")
		total := 0
		for _, cal := range foods {
			cal_int, err := strconv.Atoi(cal)
			if err != nil {
				fmt.Println(err)
				return -1
			}

			total += cal_int
		}

		cur := total
		for i, val := range max {
			if cur > val {
				max[i] = cur
				cur = val
			}
		}
	}

	total := 0
	for _, val := range max {
		total += val
	}

	return total
}

func main() {
	result_1 := part1(input)
	result_2 := part2(input)

	fmt.Println(result_1)
	fmt.Println(result_2)
}
