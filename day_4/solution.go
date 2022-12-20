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
	solution2 := part2(assignemntPairs)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %d", solution1),
		fmt.Sprintf("part1: %d", solution2),
	})
}

func part1(assignemntPairs []AssignmentPair) int {
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

func part2(assignemntPairs []AssignmentPair) int {
	overlappingPairs := 0

	for _, assignmentPair := range assignemntPairs {
		sectionRange1 := strings.Split(string(assignmentPair[0]), "-")
		sectionRange2 := strings.Split(string(assignmentPair[1]), "-")

		if areRangesOverlapping(sectionRange1, sectionRange2) {
			overlappingPairs++
		}
	}

	return overlappingPairs
}

func parseInput(lines []string) []AssignmentPair {
	assignemntPairs := []AssignmentPair{}

	// lines -> ["19-73,18-73", "1-7,8-7"]
	// line -> "19-73,18-73"
	for _, line := range lines {
		// assignments -> ["19-73","18-73"]
		assignments := strings.Split(line, ",")
		// AssignmentPair type -> ["19-73","18-73"]
		// assignmentPais -> [ ["19-73","18-73"] , ["1-7","8-7"] ]
		assignemntPairs = append(assignemntPairs, AssignmentPair{assignments[0], assignments[1]})
	}

	return assignemntPairs
}

func isOneRangeContained(sectionRange1 []string, sectionRange2 []string) bool {
	// example contained ranges
	// 2-4, 1-4
	// 3-7, 5-6

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

func areRangesOverlapping(sectionRange1 []string, sectionRange2 []string) bool {
	// example overlapping ranges, contained ranges still count
	// 2-4, 3-8
	// 3-7, 1-3
	// 2-4, 1-4
	// 3-7, 5-6

	range1Lower, _ := strconv.Atoi(sectionRange1[0])
	range1Upper, _ := strconv.Atoi(sectionRange1[1])
	range2Lower, _ := strconv.Atoi(sectionRange2[0])
	range2Upper, _ := strconv.Atoi(sectionRange2[1])

	if range1Lower <= range2Lower && range1Upper >= range2Lower {
		return true
	} else if range2Lower <= range1Lower && range2Upper >= range1Lower {
		return true
	}

	return false
}
