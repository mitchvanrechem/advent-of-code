package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, advent of code 2022!")
	inputLines := utils.ReadInputAsStrings("input.txt")

	solutions := readInventories(inputLines)
	utils.PrintSolution(solutions)
}

func readInventories(calories_entries []string) *[]string {
	calories_max := 0
	current_calories := 0
	elves := make([]int, 0)

	for _, calories_entry := range calories_entries {
		if calories_entry == "" {
			elves = append(elves, current_calories)
			//fmt.Printf("elf %d, with %d calories\n", len(elves), current_calories)

			if current_calories > calories_max {
				calories_max = current_calories
			}

			current_calories = 0
		} else {
			calories, _ := strconv.Atoi(calories_entry)
			current_calories += calories
		}

	}

	sort.Ints(elves)

	return &[]string{
		fmt.Sprintf("Largest amount of calories: %d", calories_max),
		fmt.Sprintf("Top 3 elves: %d + %d + %d", elves[len(elves)-1], elves[len(elves)-2], elves[len(elves)-3]),
		fmt.Sprintf("Top 3 elves total amount: %d", elves[len(elves)-1]+elves[len(elves)-2]+elves[len(elves)-3]),
	}
}
