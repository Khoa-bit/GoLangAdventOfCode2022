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
	height int
	left   int
	right  int
	top    int
	bottom int
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
	trim_input := strings.Trim(input, "\n")

	strRows := strings.Split(trim_input, "\n")

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
		cacheLeft := make([]int, 10)
		for left := 0; left < len(row)-1; left++ {
			curLeft := &resultGrid[rowIndex][left]

			curLeft.left = left - cacheLeft[curLeft.height]

			updateBlockedTreeCache(&cacheLeft, curLeft.height, left)
		}

		// Go from right, skip grid border (index len - 1)
		// While keep track of the tallest tree
		cacheRight := []int{colsCount - 1, colsCount - 1, colsCount - 1, colsCount - 1, colsCount - 1, colsCount - 1, colsCount - 1, colsCount - 1, colsCount - 1, colsCount - 1}
		for right := len(grid[0]) - 1; right > 0; right-- {
			curRight := &resultGrid[rowIndex][right]

			curRight.right = cacheRight[curRight.height] - right

			updateBlockedTreeCache(&cacheRight, curRight.height, right)
		}
	}

	// Vertical Scan
	for colIndex := 1; colIndex < len(grid[0])-1; colIndex++ {
		// Go from top, skip grid border (index 0)
		// While keep track of the tallest tree
		cacheTop := make([]int, 10)
		for top := 0; top < len(grid)-1; top++ {
			curTop := &resultGrid[top][colIndex]

			curTop.top = top - cacheTop[curTop.height]

			updateBlockedTreeCache(&cacheTop, curTop.height, top)
		}

		// Go from bottom, skip grid border (index len - 1)
		// While keep track of the tallest tree
		cacheBottom := []int{rowsCount - 1, rowsCount - 1, rowsCount - 1, rowsCount - 1, rowsCount - 1, rowsCount - 1, rowsCount - 1, rowsCount - 1, rowsCount - 1, rowsCount - 1}
		for bottom := len(grid) - 2; bottom > 0; bottom-- {
			curBottom := &resultGrid[bottom][colIndex]

			curBottom.bottom = cacheBottom[curBottom.height] - bottom

			updateBlockedTreeCache(&cacheBottom, curBottom.height, bottom)
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

func updateBlockedTreeCache(cache *[]int, height int, position int) *[]int {
	for h := 0; h <= height; h++ {
		(*cache)[h] = position
	}

	return cache
}

func main() {
	// result1 := part1(input)
	// fmt.Println(result1)

	result2 := part2(input)
	fmt.Println(result2)
}
