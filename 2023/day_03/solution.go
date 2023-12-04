package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	lines := utils.ReadInputAsStrings("input.txt")
	schematic := parseInput(lines)

	solution1 := part1(schematic)

	fmt.Printf("Solution to part 1: %v\n", solution1)
}

func part1(schematic [][]rune) int {

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

			if j == len(schematic[i])-1 || !unicode.IsDigit(schematic[i][j+1]) {
				if isPartNumber {
					currentNumber, _ := strconv.Atoi(currentNumberAsString)
					totalPartNumbers += currentNumber
					isPartNumber = false
					// fmt.Printf("is part number: %v ; current total: %v\n", currentNumber, totalPartNumbers)
				}

				currentNumberAsString = ""
			}
		}
	}

	return totalPartNumbers
}

func isOneSurroundingCharacterSymbol(specialCharExp *regexp.Regexp, schematic *[][]rune, i, j int, a string) bool {
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

