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

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")
	instructions := parseInput(inputLines)

	solution1 := part1(instructions)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %v", solution1),
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
