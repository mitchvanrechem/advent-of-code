package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	command        string
	signalStrength int
}

// The CRT screen is 6x40 screen
type CrtScreen [6][40]bool

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")
	instructions := parseInput(inputLines)

	solution1 := part1(instructions)
	solution2 := part2(instructions)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %v", solution1),
		fmt.Sprintf("part2: %v", solution2),
	})
}

func part1(instructions []Instruction) int {
	// Base signal strength at 1
	signalStrengthSum := 1
	nthCyclesSignalStrengthSum := 0

	for i, instruction := range instructions {
		// i is zero based indexed, cycles are one based indexed
		cycle := i + 1

		if isRelevantCycle(cycle) {
			nthCyclesSignalStrengthSum += signalStrengthSum * cycle
		}

		signalStrengthSum += instruction.signalStrength
	}

	return nthCyclesSignalStrengthSum
}

func part2(instructions []Instruction) string {
	// Sprite is 3 pixels wide and is by default at the front of the crt row
	// this variable represents the middle of the sprite
	spritePosition := 1
	crt := CrtScreen{}

	for cycle, instruction := range instructions {
		spriteIndices := []int{spritePosition - 1, spritePosition, spritePosition + 1}

		// Cycle in part 2 is zero based indexed, as it aligns with the zero
		// based indexing of the crt pixels in a row 0-39. As per cycle a pixel
		// is checked e.g. during the 4th cycle (cycle=3), the pixel index 3 is
		// checked.

		// A row has 40 pixels, at the 41st cycle (cycle=40) the first pixel of
		// the second row is drawn. Dividing and taking the modulo of cycle
		// result respectively in the row index and pixel index on that row.
		row := cycle / 40
		pixelIndex := cycle % 40

		// If drawn, pixel set to true; If not drawn , pixel set to false.
		crt[row][pixelIndex] = utils.Contains(spriteIndices, pixelIndex)

		spritePosition += instruction.signalStrength
	}

	return createCRTString(crt)
}

func createCRTString(crt CrtScreen) string {
	crtString := "\n"

	for i := range crt {
		for j := range crt[i] {
			if crt[i][j] {
				crtString += "#"
			} else {
				crtString += "."
			}
		}
		crtString += "\n"
	}

	return crtString
}

func isRelevantCycle(cycle int) bool {
	return cycle == 20 || (cycle-20)%40 == 0
}

func parseInput(lines []string) []Instruction {
	instructions := make([]Instruction, 0, len(lines))

	for _, line := range lines {
		splitLine := strings.Split(line, " ")

		switch splitLine[0] {
		case "noop":
			instructions = append(instructions, Instruction{command: "noop", signalStrength: 0})
		case "addx":
			signalStrenght, _ := strconv.Atoi(splitLine[1])

			// Because one addx command takes 2 cycles to resolve, the first
			// idle cycle is added here with signal strength value 0, this
			// simplifies the main loop.
			instructions = append(instructions, Instruction{command: "addx", signalStrength: 0})
			instructions = append(instructions, Instruction{command: "addx", signalStrength: signalStrenght})
		}

	}

	return instructions
}
