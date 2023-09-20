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

type DirNode struct {
	name     string
	size     int
	parent   *DirNode
	children []*DirNode
}

func part1(input string) int {
	cd := regexp.MustCompile(`\$ cd (.+)`)
	size := regexp.MustCompile(`(\d+) .+`)

	trim_input := strings.Trim(input, "\r\n")
	steps := strings.Split(trim_input, "\r\n")

	var root, cur *DirNode

	for i := 0; i < len(steps); i++ {
		step := steps[i]
		submatch := cd.FindStringSubmatch(step)
		sizeSubMatch := size.FindStringSubmatch(step)
		if len(submatch) > 0 {
			if root == nil {
				root = &DirNode{
					name:     submatch[1],
					size:     0,
					parent:   nil,
					children: []*DirNode{},
				}

				cur = root
				continue
			}

			if submatch[1] == ".." {
				cur = cur.parent
				continue
			}

			child := &DirNode{
				name:     submatch[1],
				size:     0,
				parent:   cur,
				children: []*DirNode{},
			}

			cur.children = append(cur.children, child)

			cur = child
		} else if len(sizeSubMatch) > 0 {
			// Extract size then cur.size += size
			sizeInt, err := strconv.Atoi(sizeSubMatch[1])
			if err != nil {
				fmt.Println(err)
				return -1
			}

			cur.size += sizeInt
		}
	}

	resultSum := 0
	bfsSumOfMax100000(root, 0, &resultSum)

	return resultSum
}

func part2(input string) int {
	cd := regexp.MustCompile(`\$ cd (.+)`)
	size := regexp.MustCompile(`(\d+) .+`)

	trim_input := strings.Trim(input, "\r\n")
	steps := strings.Split(trim_input, "\r\n")

	var root, cur *DirNode

	for i := 0; i < len(steps); i++ {
		step := steps[i]
		submatch := cd.FindStringSubmatch(step)
		sizeSubMatch := size.FindStringSubmatch(step)
		if len(submatch) > 0 {
			if root == nil {
				root = &DirNode{
					name:     submatch[1],
					size:     0,
					parent:   nil,
					children: []*DirNode{},
				}

				cur = root
				continue
			}

			if submatch[1] == ".." {
				cur = cur.parent
				continue
			}

			child := &DirNode{
				name:     submatch[1],
				size:     0,
				parent:   cur,
				children: []*DirNode{},
			}

			cur.children = append(cur.children, child)

			cur = child
		} else if len(sizeSubMatch) > 0 {
			// Extract size then cur.size += size
			sizeInt, err := strconv.Atoi(sizeSubMatch[1])
			if err != nil {
				fmt.Println(err)
				return -1
			}

			cur.size += sizeInt
		}
	}

	rootSize := bfsSize(root, 0)
	if rootSize < 40000000 {
		return 0
	}

	resultSmallestPerBranch := make([]int, 0, 20)
	bfsSmallestPerBranch(root, 0, rootSize-40000000, &resultSmallestPerBranch)

	// 2080344
	if len(resultSmallestPerBranch) == 0 {
		return -1
	}

	minBranch := resultSmallestPerBranch[0]
	for _, branch := range resultSmallestPerBranch {
		if branch < minBranch {
			minBranch = branch
		}
	}

	return minBranch
}

func bfsSize(root *DirNode, level int) int {
	sumSelfAndChildren := root.size

	for _, child := range root.children {
		sumSelfAndChildren += bfsSize(child, level+1)
	}

	return sumSelfAndChildren
}

func bfsSumOfMax100000(root *DirNode, level int, resultSum *int) int {
	sumSelfAndChildren := root.size

	for _, child := range root.children {
		sumSelfAndChildren += bfsSumOfMax100000(child, level+1, resultSum)
	}

	if sumSelfAndChildren <= 100000 {
		*resultSum += sumSelfAndChildren
	}

	return sumSelfAndChildren
}

func bfsSmallestPerBranch(root *DirNode, level int, requirement int, resultSmallestPerBranch *[]int) int {
	sumSelfAndChildren := root.size

	for _, child := range root.children {
		sumSelfAndChildren += bfsSmallestPerBranch(child, level+1, requirement, resultSmallestPerBranch)
	}

	if sumSelfAndChildren >= requirement {
		*resultSmallestPerBranch = append(*resultSmallestPerBranch, sumSelfAndChildren)
	}

	return sumSelfAndChildren
}

func main() {
	result1 := part1(input)
	result2 := part2(input)

	fmt.Println(result1)
	fmt.Println(result2)
}
