package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
)

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1(inputLines)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %v", solution1),
	})
}

func part1(inputLines []string) int {
	treeGrid := parseInput(inputLines)

	visibleTreeCoordinates := make(map[string]int)

	var tallestTop []int
	var tallestBottom []int

	for i, treeRow := range treeGrid {
		tallestLeft := treeRow[0]
		tallestRight := treeRow[len(treeRow)-1]

		for j := range treeRow {
			leftTree := treeGrid[i][j]
			topTree := treeGrid[i][j]
			rightTree := treeGrid[i][len(treeRow)-1-j]
			bottomTree := treeGrid[len(treeGrid)-1-i][j]

			leftCoordinates := fmt.Sprintf("%d-%d", i, j)
			topCoordinates := fmt.Sprintf("%d-%d", i, j)
			rightCoordinates := fmt.Sprintf("%d-%d", i, len(treeRow)-1-j)
			bottomCoordinates := fmt.Sprintf("%d-%d", len(treeGrid)-1-i, j)

			// Set initial tallest top and bottom tree for collumn with index j
			if i == 0 {
				tallestTop = append(tallestTop, treeGrid[0][j])
				tallestBottom = append(tallestBottom, treeGrid[len(treeGrid)-1][j])

				visibleTreeCoordinates[topCoordinates] = topTree
				visibleTreeCoordinates[bottomCoordinates] = bottomTree
				continue
			}

			if j == 0 {
				visibleTreeCoordinates[leftCoordinates] = leftTree
				visibleTreeCoordinates[rightCoordinates] = rightTree
				continue
			}

			// LEFT
			if leftTree > tallestLeft {
				tallestLeft = leftTree
				visibleTreeCoordinates[leftCoordinates] = leftTree
			}

			// RIGHT
			if rightTree > tallestRight {
				tallestRight = rightTree
				visibleTreeCoordinates[rightCoordinates] = rightTree
			}

			// TOP
			if topTree > tallestTop[j] {
				tallestTop[j] = topTree
				visibleTreeCoordinates[topCoordinates] = topTree
			}

			// BOTTOM
			if bottomTree > tallestBottom[j] {
				tallestBottom[j] = bottomTree
				visibleTreeCoordinates[bottomCoordinates] = bottomTree
			}

		}
	}

	amountOfVisibleTrees := len(visibleTreeCoordinates)

	return amountOfVisibleTrees
}

func parseInput(lines []string) [][]int {
	treeGrid := [][]int{}

	for _, line := range lines {
		treeRow := []int{}

		for _, letter := range line {
			tree, _ := strconv.Atoi(string(letter))
			treeRow = append(treeRow, tree)
		}

		treeGrid = append(treeGrid, treeRow)
	}

	return treeGrid
}
