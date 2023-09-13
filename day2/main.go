package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func get_result(opponent string, me string) int {
	var opponent_val int
	switch opponent {
	case "A":
		opponent_val = 1
	case "B":
		opponent_val = 2
	case "C":
		opponent_val = 3
	}

	var me_val int
	switch me {
	case "X":
		me_val = 1
	case "Y":
		me_val = 2
	case "Z":
		me_val = 3
	}

	var result int
	if opponent_val == me_val {
		result = 3
	} else if opponent_val-me_val == 1 || opponent_val-me_val == -2 {
		result = 0
	} else {
		result = 6
	}

	return me_val + result
}

func deduce_outcome(opponent string, outcome string) int {
	var opponent_val int
	switch opponent {
	case "A":
		opponent_val = 1
	case "B":
		opponent_val = 2
	case "C":
		opponent_val = 3
	}

	var outcome_val int
	var me_val int
	switch outcome {
	case "X":
		outcome_val = 0
		me_val = opponent_val - 1
		if me_val == 0 {
			me_val = 3
		}
	case "Y":
		outcome_val = 3
		me_val = opponent_val
	case "Z":
		outcome_val = 6
		me_val = opponent_val + 1
		if me_val == 4 {
			me_val = 1
		}
	}

	return me_val + outcome_val
}

func part1(input string) int {
	trim_input := strings.Trim(input, "\r\n")

	rounds := strings.Split(trim_input, "\r\n")
	score := 0
	for _, round := range rounds {
		plays := strings.Split(round, " ")
		if len(plays) != 2 {
			fmt.Println(round)
			return -1
		}

		score += get_result(plays[0], plays[1])
	}

	return score
}

func part2(input string) int {
	trim_input := strings.Trim(input, "\r\n")

	rounds := strings.Split(trim_input, "\r\n")
	score := 0
	for _, round := range rounds {
		plays := strings.Split(round, " ")
		if len(plays) != 2 {
			fmt.Println(round)
			return -1
		}

		score += deduce_outcome(plays[0], plays[1])
	}

	return score
}

func main() {
	result_1 := part1(input)
	result_2 := part2(input)

	fmt.Println(result_1)
	fmt.Println(result_2)
}
