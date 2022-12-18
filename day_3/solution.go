package main

import (
	"advent-of-code-2022/utils"
	"fmt"
)

func main() {
	inputLines := utils.ReadInput("input.txt")

	part1(inputLines)
}

func part1(lines []string) {
	for _, line := range lines {
		if line != "" {
			fmt.Println(line)
		}
	}
}
