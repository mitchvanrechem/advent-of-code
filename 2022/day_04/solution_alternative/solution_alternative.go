package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

// ["19-73","18-73"]
type AssignmentPair struct {
	first  string
	second string
}

type SectionRange struct {
	lowerBounds int
	upperBounds int
}

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	assignemntPairs := parseInput(inputLines)

	// This solution only differs from the original in typing and squashing 
	// parts 1 and 2 together, functionaly it does the exact same thing
	// This is purely me fiddling around with specific struct typing

	solution1, solution2 := part1And2(assignemntPairs)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %d", solution1),
		fmt.Sprintf("part1: %d", solution2),
	})
}

func part1And2(assignemntPairs []AssignmentPair) (int, int) {
	containedPairs := 0
	overlappingPairs := 0

	for _, assignmentPair := range assignemntPairs {
		splitAssignment1 := strings.Split(string(assignmentPair.first), "-")
		splitAssignment2 := strings.Split(string(assignmentPair.second), "-")

		range1Lower, _ := strconv.Atoi(splitAssignment1[0])
		range1Upper, _ := strconv.Atoi(splitAssignment1[1])
		range2Lower, _ := strconv.Atoi(splitAssignment2[0])
		range2Upper, _ := strconv.Atoi(splitAssignment2[1])

		sectionRange1 := SectionRange{lowerBounds: range1Lower, upperBounds: range1Upper}
		sectionRange2 := SectionRange{lowerBounds: range2Lower, upperBounds: range2Upper}

		if isOneRangeContained(sectionRange1, sectionRange2) {
			containedPairs++
		}

		if areRangesOverlapping(sectionRange1, sectionRange2) {
			overlappingPairs++
		}
	}

	return containedPairs, overlappingPairs
}

func parseInput(lines []string) []AssignmentPair {
	assignemntPairs := []AssignmentPair{}

	// lines -> ["19-73,18-73", "1-7,8-7"]
	// line -> "19-73,18-73"
	for _, line := range lines {
		// assignments -> ["19-73","18-73"]
		assignments := strings.Split(line, ",")
		// AssignmentPair type -> {"19-73","18-73"}
		// assignmentPais -> [ {"19-73","18-73"} , {"1-7","8-7"} ]
		assignemntPairs = append(assignemntPairs, AssignmentPair{assignments[0], assignments[1]})
	}

	return assignemntPairs
}

func isOneRangeContained(sr1 SectionRange, sr2 SectionRange) bool {
	// example contained ranges
	// 2-4, 1-4
	// 3-7, 5-6

	if sr1.lowerBounds <= sr2.lowerBounds && sr1.upperBounds >= sr2.upperBounds {
		return true
	} else if sr2.lowerBounds <= sr1.lowerBounds && sr2.upperBounds >= sr1.upperBounds {
		return true
	}

	return false
}

func areRangesOverlapping(sr1 SectionRange, sr2 SectionRange) bool {
	// example overlapping ranges, contained ranges still count
	// 2-4, 3-8
	// 3-7, 1-3
	// 2-4, 1-4
	// 3-7, 5-6

	if sr1.lowerBounds <= sr2.lowerBounds && sr1.upperBounds >= sr2.lowerBounds {
		return true
	} else if sr2.lowerBounds <= sr1.lowerBounds && sr2.upperBounds >= sr1.lowerBounds {
		return true
	}

	return false
}
