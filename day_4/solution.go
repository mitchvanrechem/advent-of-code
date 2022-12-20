package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

// ["19-73","18-73"]
type AssignmentPair [2]string

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	assignemntPairs := parseInput(inputLines)
	solution1 := part1(assignemntPairs)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %d", solution1),
	})
}

func part1(assignemntPairs []AssignmentPair) int {
	// example contained ranges
	// 2-4, 1-4
	// 3-7, 5-6

	containedPairs := 0

	for _, assignmentPair := range assignemntPairs {
		sectionRange1 := strings.Split(string(assignmentPair[0]), "-")
		sectionRange2 := strings.Split(string(assignmentPair[1]), "-")

		if isOneRangeContained(sectionRange1, sectionRange2) {
			containedPairs++
		}
	}

	return containedPairs
}

func parseInput(lines []string) []AssignmentPair {
	assignemntPairs := []AssignmentPair{}

	// lines -> ["19-73,18-73", "1-7,8-7"]
	// line -> "19-73,18-73"
	for _, line := range lines {
		// assignments -> ["19-73","18-73"]
		assignments := strings.Split(line, ",")
		// assignmentPair -> ["19-73","18-73"]
		// assignmentPais -> [ ["19-73","18-73"] , ["1-7","8-7"] ]
		assignemntPairs = append(assignemntPairs, AssignmentPair{assignments[0], assignments[1]})
	}

	return assignemntPairs
}

func isOneRangeContained(sectionRange1 []string, sectionRange2 []string) bool {
	range1Lower, _ := strconv.Atoi(sectionRange1[0])
	range1Upper, _ := strconv.Atoi(sectionRange1[1])
	range2Lower, _ := strconv.Atoi(sectionRange2[0])
	range2Upper, _ := strconv.Atoi(sectionRange2[1])

	if range1Lower <= range2Lower && range1Upper >= range2Upper {
		return true
	} else if range2Lower <= range1Lower && range2Upper >= range1Upper {
		return true
	}

	return false
}
