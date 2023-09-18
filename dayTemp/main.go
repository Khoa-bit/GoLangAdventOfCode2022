package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getRange(elf string) ([]int, error) {
	elfStrRange := strings.Split(elf, "-")

	lower, err := strconv.Atoi(elfStrRange[0])
	if err != nil {
		return nil, err
	}

	upper, err := strconv.Atoi(elfStrRange[1])
	if err != nil {
		return nil, err
	}

	return []int{lower, upper}, nil
}

func part1(input string) (int, error) {
	trimInput := strings.Trim(input, "\r\n")

	fullCount := 0
	pairs := strings.Split(trimInput, "\r\n")
	for _, pair := range pairs {
		elfs := strings.Split(pair, ",")
		elf1, err := getRange(elfs[0])
		if err != nil {
			return -1, err
		}

		elf2, err := getRange(elfs[1])
		if err != nil {
			return -1, err
		}

		if (elf1[0] <= elf2[0] && elf1[1] >= elf2[1]) || (elf2[0] <= elf1[0] && elf2[1] >= elf1[1]) {
			fullCount += 1
		}
	}

	return fullCount, nil
}

func part2(input string) (int, error) {
	trimInput := strings.Trim(input, "\r\n")

	fullCount := 0
	pairs := strings.Split(trimInput, "\r\n")
	for _, pair := range pairs {
		elfs := strings.Split(pair, ",")
		elf1, err := getRange(elfs[0])
		if err != nil {
			return -1, err
		}

		elf2, err := getRange(elfs[1])
		if err != nil {
			return -1, err
		}

		if elf2[1] < elf1[0] || elf1[1] < elf2[0] {
			continue
		}
		fullCount += 1
	}

	return fullCount, nil
}

func main() {
	result1, err := part1(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result1)

	result2, err := part2(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result2)
}
