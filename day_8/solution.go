package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
)

type TreeGrid [][]int

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")
	treeGrid := parseInput(inputLines)

	solution1 := part1(treeGrid)
	solution2 := part2(treeGrid)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %v", solution1),
		fmt.Sprintf("part2: %v", solution2),
	})
}

func part1(treeGrid TreeGrid) int {
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

func part2(treeGrid TreeGrid) int {
	highestScenicScore := 0

	for i, treeRow := range treeGrid {
		for j := range treeRow {
			leftScore := treeGrid.getLeftScenicScore(i, j)
			rightScore := treeGrid.getRightScenicScore(i, j)
			topScore := treeGrid.getTopScenicScore(i, j)
			bottomScore := treeGrid.getBottomScenicScore(i, j)

			currentScenicScore := leftScore * rightScore * topScore * bottomScore

			if highestScenicScore < currentScenicScore {
				highestScenicScore = currentScenicScore
			}
		}
	}

	return highestScenicScore
}

func (treeGrid *TreeGrid) getLeftScenicScore(xCoordinate, yCoordinate int) int {
	if xCoordinate == 0 {
		return 0
	}

	currentTreeHeight := (*treeGrid)[xCoordinate][yCoordinate]
	scenicScore := 0

	for relativeXCoordinate := xCoordinate - 1; relativeXCoordinate >= 0; relativeXCoordinate-- {
		scenicScore += 1

		if currentTreeHeight <= (*treeGrid)[relativeXCoordinate][yCoordinate] {
			return scenicScore
		}
	}

	return scenicScore
}

func (treeGrid *TreeGrid) getRightScenicScore(xCoordinate, yCoordinate int) int {
	rightBorderX := len((*treeGrid)[0]) - 1

	if xCoordinate == rightBorderX {
		return 0
	}

	currentTreeHeight := (*treeGrid)[xCoordinate][yCoordinate]
	scenicScore := 0

	for relativeXCoordinate := xCoordinate + 1; relativeXCoordinate <= rightBorderX; relativeXCoordinate++ {
		scenicScore += 1

		if currentTreeHeight <= (*treeGrid)[relativeXCoordinate][yCoordinate] {
			return scenicScore
		}
	}

	return scenicScore
}

func (treeGrid *TreeGrid) getTopScenicScore(xCoordinate, yCoordinate int) int {
	if yCoordinate == 0 {
		return 0
	}

	currentTreeHeight := (*treeGrid)[xCoordinate][yCoordinate]
	scenicScore := 0

	for relativeYCoordinate := yCoordinate - 1; relativeYCoordinate >= 0; relativeYCoordinate-- {
		scenicScore += 1

		if currentTreeHeight <= (*treeGrid)[xCoordinate][relativeYCoordinate] {
			return scenicScore
		}
	}

	return scenicScore
}

func (treeGrid *TreeGrid) getBottomScenicScore(xCoordinate, yCoordinate int) int {
	bottomBorderY := len(*treeGrid) - 1

	if yCoordinate == bottomBorderY {
		return 0
	}

	currentTreeHeight := (*treeGrid)[xCoordinate][yCoordinate]
	scenicScore := 0

	for relativeYCoordinate := yCoordinate + 1; relativeYCoordinate <= bottomBorderY; relativeYCoordinate++ {
		scenicScore += 1

		if currentTreeHeight <= (*treeGrid)[xCoordinate][relativeYCoordinate] {
			return scenicScore
		}
	}

	return scenicScore
}

func parseInput(lines []string) TreeGrid {
	treeGrid := TreeGrid{}

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
