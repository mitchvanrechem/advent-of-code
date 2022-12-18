package main

import (
	"advent-of-code-2022/utils"
	"fmt"
)

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1(inputLines)

	utils.PrintSolution(&[]string{fmt.Sprintf("part1: %d", solution1)})
}

func part1(rucksacks []string) int {
	priorityTotal := 0

	for _, rucksack := range rucksacks {
		item := getMisplacedItem(rucksack)
		priorityTotal += determinePriority(item)
	}

	return priorityTotal
}

func getMisplacedItem(rucksack string) rune {
	compartment1 := rucksack[:(len(rucksack) / 2)]
	compartment2 := rucksack[len(rucksack)/2:]

	for _, item := range compartment2 {

		if utils.Contains([]rune(compartment1), item) {
			return item
		}
	}

	panic("no mismatched item found!")
}

func determinePriority(item rune) int {
	// ASCII code a == 97
	// ASCII code A == 65
	// priority a == 1 , z == 26
	// priority A == 27 , Z == 52

	code := int(item)

	// Check if item is upper or lower
	if byte(item) >= byte('a') {
		return code - 96
	}

	return code - 38
}
