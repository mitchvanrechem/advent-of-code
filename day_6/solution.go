package main

import (
	"advent-of-code-2022/utils"
	"fmt"
)

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1and2(inputLines[0], 4)
	solution2 := part1and2(inputLines[0], 14)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %s", solution1),
		fmt.Sprintf("part2: %s", solution2),
	})
}

func part1and2(inputLine string, packetSize int) string {
	for i := 0; i+packetSize <= len(inputLine); i++ {
		subString := inputLine[i : i+packetSize]
		marker := []rune{}
		unique := true

		for _, subChar := range subString {
			if utils.Contains(marker, subChar) {
				unique = false
				break
			}

			marker = append(marker, subChar)
		}

		if unique {
			return fmt.Sprint(i + packetSize)
		}
	}

	panic("No marker found!")
}
