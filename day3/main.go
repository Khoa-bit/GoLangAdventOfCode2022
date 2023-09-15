package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func findMatch(rucksack string) byte {
	half_ptr := len(rucksack) / 2
	for i := 0; i < half_ptr; i++ {
		for j := half_ptr; j < len(rucksack); j++ {
			if rucksack[i] == rucksack[j] {
				return rucksack[i]
			}
		}
	}
	return 0
}

func convertScore(char byte) byte {
	var score byte
	switch {
	case 97 <= char && char <= 122:
		score = char - 96
	case 65 <= char && char <= 90:
		score = char - 64 + 26
	}
	return score
}

func part1(input string) int {
	trim_input := strings.Trim(input, "\r\n")
	rucksacks := strings.Split(trim_input, "\r\n")
	total := 0
	for _, rucksack := range rucksacks {
		total += int(convertScore(findMatch(rucksack)))
	}

	return total
}

func part2(input string) int {
	trim_input := strings.Trim(input, "\r\n")
	rucksacks := strings.Split(trim_input, "\r\n")
	total := 0
	for i := 0; i < len(rucksacks); i += 3 {
		counter := make([]int, 52)

		anchor := i%3 + i
		for _, item := range rucksacks[anchor] {
			index := convertScore(byte(item)) - 1
			counter[index] = 1
		}

		for _, item := range rucksacks[anchor+1] {
			index := convertScore(byte(item)) - 1
			if counter[index] == 1 {
				counter[index] = 2
			}
		}

		for _, item := range rucksacks[anchor+2] {
			index := convertScore(byte(item)) - 1
			if counter[index] == 2 {
				total += int(convertScore(byte(item)))
				break
			}
		}
	}

	return total
}

func main() {
	result_1 := part1(input)
	result_2 := part2(input)

	fmt.Println(result_1)
	fmt.Println(result_2)
}
