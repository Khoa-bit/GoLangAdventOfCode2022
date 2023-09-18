package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string) string {
	stacks := [][]rune{
		{'W', 'R', 'T', 'G'},
		{'W', 'V', 'S', 'M', 'P', 'H', 'C', 'G'},
		{'M', 'G', 'S', 'T', 'L', 'C'},
		{'F', 'R', 'W', 'M', 'D', 'H', 'J'},
		{'J', 'F', 'W', 'S', 'H', 'L', 'Q', 'P'},
		{'S', 'M', 'F', 'N', 'D', 'J', 'P'},
		{'J', 'S', 'C', 'G', 'F', 'D', 'B', 'Z'},
		{'B', 'T', 'R'},
		{'C', 'L', 'W', 'N', 'H'},
	}

	// Reverse stack
	for _, stack := range stacks {
		start := 0
		end := len(stack) - 1
		for start < end {
			stack[start], stack[end] = stack[end], stack[start]
			start++
			end--
		}
	}

	trim_input := strings.Trim(input, "\r\n")

	steps := strings.Split(trim_input, "\n")

	re := regexp.MustCompile(`\d+`)
	for _, step := range steps {
		stepValue := re.FindAllString(step, 3)
		amount, err := strconv.Atoi(stepValue[0])
		if err != nil {
			fmt.Println(err)
			return ""
		}
		from, err := strconv.Atoi(stepValue[1])
		if err != nil {
			fmt.Println(err)
			return ""
		}
		to, err := strconv.Atoi(stepValue[2])
		if err != nil {
			fmt.Println(err)
			return ""
		}

		from--
		to--

		for i := 0; i < amount; i++ {
			// pop from stackFrom and append to stackTo
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	result := ""
	for _, stack := range stacks {
		result += fmt.Sprint(string(stack[len(stack)-1]))
	}

	return result
}

func part2(input string) string {
	stacks := [][]rune{
		{'W', 'R', 'T', 'G'},
		{'W', 'V', 'S', 'M', 'P', 'H', 'C', 'G'},
		{'M', 'G', 'S', 'T', 'L', 'C'},
		{'F', 'R', 'W', 'M', 'D', 'H', 'J'},
		{'J', 'F', 'W', 'S', 'H', 'L', 'Q', 'P'},
		{'S', 'M', 'F', 'N', 'D', 'J', 'P'},
		{'J', 'S', 'C', 'G', 'F', 'D', 'B', 'Z'},
		{'B', 'T', 'R'},
		{'C', 'L', 'W', 'N', 'H'},
	}

	// Reverse stack
	for _, stack := range stacks {
		start := 0
		end := len(stack) - 1
		for start < end {
			stack[start], stack[end] = stack[end], stack[start]
			start++
			end--
		}
	}

	trim_input := strings.Trim(input, "\r\n")

	steps := strings.Split(trim_input, "\n")

	re := regexp.MustCompile(`\d+`)
	for _, step := range steps {
		stepValue := re.FindAllString(step, 3)
		amount, err := strconv.Atoi(stepValue[0])
		if err != nil {
			fmt.Println(err)
			return ""
		}
		from, err := strconv.Atoi(stepValue[1])
		if err != nil {
			fmt.Println(err)
			return ""
		}
		to, err := strconv.Atoi(stepValue[2])
		if err != nil {
			fmt.Println(err)
			return ""
		}

		from--
		to--

		for i := 0; i < amount; i++ {
			// pop from stackFrom and append to stackTo
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-(amount-i)])
		}

		stacks[from] = stacks[from][:len(stacks[from])-amount]

	}

	result := ""
	for _, stack := range stacks {
		result += fmt.Sprint(string(stack[len(stack)-1]))
	}

	return result
}

func main() {
	result1 := part1(input)
	result2 := part2(input)

	fmt.Println(result1)
	fmt.Println(result2)
}
