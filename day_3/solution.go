package main

import (
	"advent-of-code-2022/utils"
	"fmt"
)

type Rucksack string

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	ruckSacks := parseInput(inputLines)
	solution1 := part1(ruckSacks)
	solution2 := part2(ruckSacks)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %d", solution1),
		fmt.Sprintf("part2: %d", solution2),
	})
}

func part1(rucksacks []Rucksack) int {
	priorityTotal := 0

	for _, rucksack := range rucksacks {
		item := getMisplacedItem(rucksack)
		priorityTotal += determinePriority(item)
	}

	return priorityTotal
}

func part2(rucksacks []Rucksack) int {
	priorityTotal := 0

	for i := 0; i < len(rucksacks); i += 3 {
		item := getBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		fmt.Println(string(item))
		priorityTotal += determinePriority(item)
	}

	return priorityTotal
}

func parseInput(lines []string) []Rucksack {
	rucksacks := []Rucksack{}
	for _, line := range lines {
		rucksacks = append(rucksacks, Rucksack(line))
	}

	return rucksacks
}

func getBadge(sack1 Rucksack, sack2 Rucksack, sack3 Rucksack) rune {
	for _, item := range sack1 {
		if utils.Contains([]rune(sack2), item) {
			if utils.Contains([]rune(sack3), item) {
				return item
			}
		}
	}

	panic("No badge was found in either of the 3 rucksacks")
}

func getMisplacedItem(rucksack Rucksack) rune {
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
