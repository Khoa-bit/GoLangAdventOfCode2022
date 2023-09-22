package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type TreeView struct {
	height   int
	left     int
	right    int
	top      int
	bottom   int
	leftEq   int
	rightEq  int
	topEq    int
	bottomEq int
}

func part1(input string) int {
	trim_input := strings.Trim(input, "\r\n")

	strRows := strings.Split(trim_input, "\r\n")

	rowsCount := len(strRows)
	colsCount := len(strRows[0])
	grid := make([][]int, rowsCount)

	// Parse and build grid
	for row, strRow := range strRows {
		grid[row] = make([]int, colsCount)
		for col, runeHeight := range strRow {
			height, err := strconv.Atoi(string(runeHeight))
			if err != nil {
				fmt.Println(err)
				return -1
			}
			grid[row][col] = height
		}
	}

	visitedSet := make(map[string]struct{}, 3000)

	// Horizontal Scan
	for rowIndex := 1; rowIndex < len(grid)-1; rowIndex++ {
		row := grid[rowIndex]

		// Go from left, skip grid border (index 0)
		// While keep track of the tallest tree
		maxLeft := row[0]
		for left := 1; left < len(row)-1; left++ {
			if row[left] > maxLeft {
				visitedSet[fmt.Sprintf("%d-%d", rowIndex, left)] = struct{}{}
				maxLeft = row[left]
			}
			if maxLeft == 9 {
				break
			}
		}

		// Go from right, skip grid border (index len - 1)
		// While keep track of the tallest tree
		maxRight := row[len(row)-1]
		for right := len(row) - 2; right > 0; right-- {
			if row[right] > maxRight {
				visitedSet[fmt.Sprintf("%d-%d", rowIndex, right)] = struct{}{}
				maxRight = row[right]
			}

			if maxRight == 9 {
				break
			}
		}
	}

	// Vertical Scan
	for colIndex := 1; colIndex < len(grid[0])-1; colIndex++ {
		// Go from top, skip grid border (index 0)
		// While keep track of the tallest tree
		maxTop := grid[0][colIndex]
		for top := 1; top < len(grid[0])-1; top++ {
			if grid[top][colIndex] > maxTop {
				visitedSet[fmt.Sprintf("%d-%d", top, colIndex)] = struct{}{}
				maxTop = grid[top][colIndex]
			}

			if maxTop == 9 {
				break
			}
		}

		// Go from bottom, skip grid border (index len - 1)
		// While keep track of the tallest tree
		maxBottom := grid[len(grid[0])-1][colIndex]
		for bottom := len(grid[0]) - 2; bottom > 0; bottom-- {
			if grid[bottom][colIndex] > maxBottom {
				visitedSet[fmt.Sprintf("%d-%d", bottom, colIndex)] = struct{}{}
				maxBottom = grid[bottom][colIndex]
			}

			if maxBottom == 9 {
				break
			}
		}
	}

	return len(visitedSet) + rowsCount*2 + colsCount*2 - 4
}

func part2(input string) int {
	trim_input := strings.Trim(input, "\r\n")

	strRows := strings.Split(trim_input, "\r\n")

	rowsCount := len(strRows)
	colsCount := len(strRows[0])

	// Parse and build grid
	grid := make([][]int, rowsCount)
	for row, strRow := range strRows {
		grid[row] = make([]int, colsCount)
		for col, runeHeight := range strRow {
			height, err := strconv.Atoi(string(runeHeight))
			if err != nil {
				fmt.Println(err)
				return -1
			}
			grid[row][col] = height
		}
	}

	// Parse and build result grid
	resultGrid := make([][]TreeView, rowsCount)
	for row, strRow := range strRows {
		resultGrid[row] = make([]TreeView, colsCount)
		for col, runeHeight := range strRow {
			height, err := strconv.Atoi(string(runeHeight))
			if err != nil {
				fmt.Println(err)
				return -1
			}
			resultGrid[row][col] = TreeView{height: height}
		}
	}

	// Horizontal Scan
	for rowIndex := 1; rowIndex < len(grid)-1; rowIndex++ {
		row := grid[rowIndex]

		// Go from left, skip grid border (index 0)
		// While keep track of the tallest tree
		prevLeft := &resultGrid[rowIndex][0]
		for left := 1; left < len(row)-1; left++ {
			curLeft := &resultGrid[rowIndex][left]

			if prevLeft.height <= curLeft.height {
				curLeft.leftEq = prevLeft.leftEq + 1
			} else {
				curLeft.leftEq = 1
			}

			if prevLeft.height < curLeft.height {
				curLeft.left = curLeft.leftEq
			} else {
				curLeft.left = 1
			}

			prevLeft = curLeft
		}

		// Go from right, skip grid border (index len - 1)
		// While keep track of the tallest tree
		prevRight := &resultGrid[rowIndex][len(grid[0])-1]
		for right := len(grid[0]) - 2; right > 0; right-- {
			curRight := &resultGrid[rowIndex][right]

			if prevRight.height <= curRight.height {
				curRight.rightEq = prevRight.rightEq + 1
			} else {
				curRight.rightEq = 1
			}

			if prevRight.height < curRight.height {
				curRight.right = curRight.rightEq
			} else {
				curRight.right = 1
			}

			prevRight = curRight
		}
	}

	// Vertical Scan
	for colIndex := 1; colIndex < len(grid[0])-1; colIndex++ {
		// Go from top, skip grid border (index 0)
		// While keep track of the tallest tree
		prevTop := &resultGrid[0][colIndex]
		for top := 1; top < len(grid)-1; top++ {
			curTop := &resultGrid[top][colIndex]

			if prevTop.height < curTop.height {
				curTop.top = prevTop.top + 1
			} else {
				curTop.top = 1
			}

			prevTop = curTop
		}

		// Go from bottom, skip grid border (index len - 1)
		// While keep track of the tallest tree
		prevBottom := &resultGrid[len(grid)-1][colIndex]
		for bottom := len(grid) - 2; bottom > 0; bottom-- {
			curBottom := &resultGrid[bottom][colIndex]

			if prevBottom.height < curBottom.height {
				curBottom.bottom = prevBottom.bottom + 1
			} else {
				curBottom.bottom = 1
			}

			prevBottom = curBottom
		}
	}

	highestScenicScore := 0
	highestScenicLocation := ""
	for row, rowTreeViews := range resultGrid {
		for col, treeView := range rowTreeViews {
			scenicScore := treeView.left * treeView.right * treeView.top * treeView.bottom
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
				highestScenicLocation = fmt.Sprintf("%d-%d", row, col)
			}
		}
	}

	fmt.Println("Location: " + highestScenicLocation)
	return highestScenicScore
}

func main() {
	// result1 := part1(input)
	// fmt.Println(result1)

	result2 := part2(input)
	fmt.Println(result2)
}
