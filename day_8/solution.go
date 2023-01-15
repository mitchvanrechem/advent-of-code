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

	// Arrays to keep track of the tallest tree in each collumn, of both top
	// and bottom perspective 
	var tallestTop []int
	var tallestBottom []int

	for i, treeRow := range treeGrid {
		tallestLeft := treeRow[0]
		tallestRight := treeRow[len(treeRow)-1]

		for j := range treeRow {
			leftBorderTree := treeGrid[i][j]
			topBorderTree := treeGrid[i][j]
			rightBorderTree := treeGrid[i][len(treeRow)-1-j]
			bottomBorderTree := treeGrid[len(treeGrid)-1-i][j]

			leftCoordinates := fmt.Sprintf("%d-%d", i, j)
			topCoordinates := fmt.Sprintf("%d-%d", i, j)
			rightCoordinates := fmt.Sprintf("%d-%d", i, len(treeRow)-1-j)
			bottomCoordinates := fmt.Sprintf("%d-%d", len(treeGrid)-1-i, j)

			// Handle 'edge' case of every single tree on the top and bottom
			// border being visible by default
			if i == 0 {
				// Set initial tallest top and bottom perspective tree for 
				// collumn with index j.
				tallestTop = append(tallestTop, treeGrid[0][j])
				tallestBottom = append(tallestBottom, treeGrid[len(treeGrid)-1][j])

				visibleTreeCoordinates[topCoordinates] = topBorderTree
				visibleTreeCoordinates[bottomCoordinates] = bottomBorderTree
				continue
			}

			// Handle 'edge' case of every single tree on the left and right
			// border being visible by default
			if j == 0 {
				visibleTreeCoordinates[leftCoordinates] = leftBorderTree
				visibleTreeCoordinates[rightCoordinates] = rightBorderTree
				continue
			}

			// LEFT
			if leftBorderTree > tallestLeft {
				tallestLeft = leftBorderTree
				visibleTreeCoordinates[leftCoordinates] = leftBorderTree
			}

			// RIGHT
			if rightBorderTree > tallestRight {
				tallestRight = rightBorderTree
				visibleTreeCoordinates[rightCoordinates] = rightBorderTree
			}

			// TOP
			if topBorderTree > tallestTop[j] {
				tallestTop[j] = topBorderTree
				visibleTreeCoordinates[topCoordinates] = topBorderTree
			}

			// BOTTOM
			if bottomBorderTree > tallestBottom[j] {
				tallestBottom[j] = bottomBorderTree
				visibleTreeCoordinates[bottomCoordinates] = bottomBorderTree
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
