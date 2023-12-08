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

func part1(input string) int {
	trim_input := strings.Trim(input, "\r\n")

	strRows := strings.Split(trim_input, "\r\n")
	rowsCount := len(strRows)

	addxRegex := regexp.MustCompile(`addx (.+)`)
	atOp := 0
	wait := 0
	action := 0
	regsiterX := 1
	inspectPoint := 20
	signalStrength := 0

	for cycle := 1; cycle <= 220 && atOp <= rowsCount; cycle++ {
		if cycle == inspectPoint {
			fmt.Println(cycle, regsiterX, inspectPoint*regsiterX)
			signalStrength += inspectPoint * regsiterX
			fmt.Println("signalStrength", signalStrength)
			inspectPoint += 40
		}

		if wait > 0 {
			wait -= 1
			if wait == 0 {
				// Flush action after the current cycle
				regsiterX += action
			}
			continue
		}

		strRow := strRows[atOp]
		submatch := addxRegex.FindStringSubmatch(strRow)
		if len(submatch) == 0 {
			// noop
			wait = 0
			action = 0
		} else {
			// addx
			updateValue, err := strconv.Atoi(string(submatch[1]))
			if err != nil {
				fmt.Println(err)
				return -1
			}
			wait = 1
			action = updateValue
		}

		atOp += 1
	}
	return signalStrength
}

func part2(input string) int {
	trim_input := strings.Trim(input, "\r\n")

	strRows := strings.Split(trim_input, "\r\n")
	rowsCount := len(strRows)

	addxRegex := regexp.MustCompile(`addx (.+)`)
	atOp := 0
	wait := 0
	action := 0
	regsiterX := 1
	inspectPoint := 20
	signalStrength := 0

	width := 40
	height := 6
	CRTScreen := make([][]bool, height)
	for row := 0; row < height; row++ {
		CRTScreen[row] = make([]bool, width)
	}

	for cycle := 1; cycle <= 240 && atOp <= rowsCount; cycle++ {
		if cycle == inspectPoint {
			fmt.Println(cycle, regsiterX, inspectPoint*regsiterX)
			signalStrength += inspectPoint * regsiterX
			fmt.Println("signalStrength", signalStrength)
			inspectPoint += 40
		}

		pixelAtRow := (cycle - 1) / width
		pixel := (cycle - 1) - pixelAtRow*width
		if pixel <= regsiterX+1 && pixel >= regsiterX-1 {
			CRTScreen[pixelAtRow][pixel] = true
		}

		if wait > 0 {
			wait -= 1
			if wait == 0 {
				// Flush action after the current cycle
				regsiterX += action
			}
			continue
		}

		strRow := strRows[atOp]
		submatch := addxRegex.FindStringSubmatch(strRow)
		if len(submatch) == 0 {
			// noop
			wait = 0
			action = 0
		} else {
			// addx
			updateValue, err := strconv.Atoi(string(submatch[1]))
			if err != nil {
				fmt.Println(err)
				return -1
			}
			wait = 1
			action = updateValue
		}

		atOp += 1
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if CRTScreen[row][col] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return signalStrength
}

func main() {
	// result1 := part1(input)
	// fmt.Println(result1)

	result2 := part2(input)
	fmt.Println(result2)
}
