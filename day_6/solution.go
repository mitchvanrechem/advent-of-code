package main

import (
	"advent-of-code-2022/utils"
	"fmt"
)

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1(inputLines[0])

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %s", solution1),
	})
}

func part1(inputLine string) string {
	for i := 0; i+4 <= len(inputLine); i++ {
		subString := inputLine[i : i+4]
		marker := []rune{}
		unique := true

		for _, subChar := range subString {
			if utils.Contains(marker, subChar) {
				unique = false
			}

			marker = append(marker, subChar)
		}

		if unique {
			return fmt.Sprint(i + 4)
		}
	}

	panic("No marker found!")
}
