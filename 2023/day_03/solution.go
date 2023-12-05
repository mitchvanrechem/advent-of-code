package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	lines := utils.ReadInputAsStrings("example_input.txt")
	schematic := parseInput(lines)

	solution1 := part1(schematic)
	solution2 := part2(schematic)

	fmt.Printf("Solution to part 1: %v\n", solution1)
	fmt.Printf("Solution to part 2: %v\n", solution2)
}

func parseInput(lines []string) [][]rune {
	schematic := make([][]rune, len(lines))

	for i := 0; i < len(lines); i++ {
		schematicLine := make([]rune, len(lines[i]))

		for j := 0; j < len(lines[i]); j++ {
			schematicLine[j] = rune(lines[i][j])
		}

		schematic[i] = schematicLine
	}

	return schematic
}

type Schematic [][]rune

func part1(schematic Schematic) int {

	specialCharacterExp := regexp.MustCompile(`([^0-9\.])`)

	currentNumberAsString := ""
	var totalPartNumbers int = 0
	var isPartNumber bool = false

	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			currentChar := schematic[i][j]

			if !unicode.IsDigit(currentChar) {
				continue
			}

			if unicode.IsDigit(currentChar) {
				currentNumberAsString += string(currentChar)

				if !isPartNumber {
					isPartNumber = isOneSurroundingCharacterSymbol(specialCharacterExp, &schematic, i, j, string(currentChar))
				}
			}

			// Either the number is at the end of the row or the next character is not a digit
			// In both cases the number is fully discovered, the build string can be parsed and added to the total
			if j == len(schematic[i])-1 || !unicode.IsDigit(schematic[i][j+1]) {
				if isPartNumber {
					currentNumber, _ := strconv.Atoi(currentNumberAsString)
					totalPartNumbers += currentNumber
					isPartNumber = false
				}

				currentNumberAsString = ""
			}
		}
	}

	return totalPartNumbers
}

func isOneSurroundingCharacterSymbol(specialCharExp *regexp.Regexp, schematic *Schematic, i, j int, a string) bool {
	upperBound := max(i-1, 0)
	lowerBound := min(i+1, len(*schematic)-1)
	leftBound := max(j-1, 0)
	rightBound := min(j+1, len((*schematic)[i])-1)

	surroundingCharacters := []rune{
		(*schematic)[upperBound][leftBound],  // top left
		(*schematic)[upperBound][j],          // top
		(*schematic)[upperBound][rightBound], // top right
		(*schematic)[i][rightBound],          // right
		(*schematic)[lowerBound][rightBound], // bottom right
		(*schematic)[lowerBound][j],          // bottom
		(*schematic)[lowerBound][leftBound],  // bottom left
		(*schematic)[i][leftBound],           // left
	}

	for _, character := range surroundingCharacters {
		if specialCharExp.MatchString(string(character)) {
			return true
		}
	}

	return false
}

type Coord struct {
	x int
	y int
}

type CoordMap map[Coord]int

func part2(schematic Schematic) int {
	gearRatioSum := 0
	isGear := false
	currentNumberAsString := ""
	var currentGearCoord Coord

	// Map of the asterisk coordinates to the first number that makes up a gear ratio
	coordMap := make(CoordMap)

	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {

			currentChar := schematic[i][j]

			if !unicode.IsDigit(currentChar) {
				continue
			}

			if unicode.IsDigit(currentChar) {
				currentNumberAsString += string(currentChar)

				if !isGear {
					coord := getSurroundingAsteriskCoord(&schematic, i, j, string(currentChar))
					if coord != (Coord{}) {
						isGear = true
						currentGearCoord = coord
					}
				}

			}

			// Either the number is at the end of the row or the next character is not a digit
			// In both cases the number is fully discovered, the build string can be parsed and used to calculate the gear ratio
			if j == len(schematic[i])-1 || !unicode.IsDigit(schematic[i][j+1]) {

				if isGear {
					currentNumber, _ := strconv.Atoi(currentNumberAsString)

					// Check if the asterisk has been spotted before
					// If the coord of the asterisk is not in the map, add the current number which makes up on part of the gear ratio
					// If the coord of the asterisk is in the map, use the number that is the value to mupltiply with the current number
					// sharing the same coord
					if coordMap[currentGearCoord] == 0 {
						coordMap[currentGearCoord] = currentNumber
					} else {
						gearRatio := currentNumber * coordMap[currentGearCoord]
						gearRatioSum += gearRatio
					}

					isGear = false
				}

				currentNumberAsString = ""
			}
		}
	}

	return gearRatioSum
}

func getSurroundingAsteriskCoord(schematic *Schematic, i, j int, a string) Coord {
	// Deal with the edges of the matrix, to not have out of range exceptions
	// currently a naive approach, that will produce duplicate coords and unnecessary repeated checks
	upperBound := max(i-1, 0)
	lowerBound := min(i+1, len(*schematic)-1)
	leftBound := max(j-1, 0)
	rightBound := min(j+1, len((*schematic)[i])-1)

	surroundingCharacters := []Coord{
		{y: upperBound, x: leftBound},  // top left
		{y: upperBound, x: j},          // top
		{y: upperBound, x: rightBound}, // top right
		{y: i, x: rightBound},          // right
		{y: lowerBound, x: rightBound}, // bottom right
		{y: lowerBound, x: j},          // bottom
		{y: lowerBound, x: leftBound},  // bottom left
		{y: i, x: leftBound},           // left
	}

	for _, coord := range surroundingCharacters {
		if (*schematic)[coord.y][coord.x] == '*' {
			return coord
		}
	}

	return Coord{}
}
