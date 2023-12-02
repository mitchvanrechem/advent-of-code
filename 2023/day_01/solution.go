package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	lines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1(lines)
	solution2 := part2(lines)

	fmt.Printf("Solution to part 1: %v\n", solution1)
	fmt.Printf("Solution to part 2: %v\n", solution2)
}

func part1(lines []string) int {
	totalCalibrationValues := 0

	for _, line := range lines {
		characters := []rune(line)

		var firstDigit, lastDigit rune

		for i := 0; i < len(characters); i++ {
			char := characters[i]

			if unicode.IsDigit(char) {
				firstDigit = char
				break
			}
		}

		for i := len(characters) - 1; i >= 0; i-- {
			char := characters[i]

			if unicode.IsDigit(char) {
				lastDigit = char
				break
			}
		}

		calibrationValueAsString := string(firstDigit) + string(lastDigit)

		calibrationValue, err := strconv.Atoi(calibrationValueAsString)
		if err != nil {
			fmt.Printf("unable to convert strings to int: %v", calibrationValueAsString)
			panic(err)
		}

		totalCalibrationValues += calibrationValue
	}

	return totalCalibrationValues
}

var spelledOutDigits = [9]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func part2(lines []string) int {
	totalCalibrationValues := 0

	for _, line := range lines {

		characters := []rune(line)

		var firstDigit, lastDigit int
		var firstDigitIndex, lastDigitIndex int

		for i := 0; i < len(characters); i++ {
			char := characters[i]

			if unicode.IsDigit(char) {
				firstDigit, _ = strconv.Atoi(string(char))
				firstDigitIndex = i
				break
			}
		}

		for i := len(characters) - 1; i >= 0; i-- {
			char := characters[i]

			if unicode.IsDigit(char) {
				lastDigit, _ = strconv.Atoi(string(char))
				lastDigitIndex = i
				break
			}
		}


		// Setting these to the length of the line and 0 respectively to have 
		// the outer bounds on both , ensuring that on the first substring match 
		// the prev indices are always set, if all other conditions are also met
		var currFirstIndex, currLastIndex int = len(line), 0 

		for i, spelledOutDigit := range spelledOutDigits {
			firstIndex := strings.Index(line, spelledOutDigit)
			lastIndex := strings.LastIndex(line, spelledOutDigit)

			if firstIndex != -1 && firstIndex < currFirstIndex && firstIndex < firstDigitIndex {
				firstDigit = i + 1
				currFirstIndex = firstIndex
			}

			if lastIndex != -1 && lastIndex > currLastIndex && lastIndex > lastDigitIndex {
				lastDigit = i + 1
				currLastIndex = lastIndex
			}
		}

		calibrationValueAsString := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
		calibrationValue, err := strconv.Atoi(calibrationValueAsString)
		if err != nil {
			fmt.Printf("unable to convert strings to int: %v", calibrationValueAsString)
			panic(err)
		}

		totalCalibrationValues += calibrationValue
	}

	return totalCalibrationValues
}
