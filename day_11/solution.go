package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"
)

type Monkey struct {
	items              []int
	operation          func(int) int
	divisionTestNumber int
	positiveTestMonkey int
	negativeTestMonkey int
}

func main() {
	// Using hard coded monkeys
	solution1 := part1(getMonkeys())
	solution2 := part2(getMonkeys())

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %v", solution1),
		fmt.Sprintf("part2: %v", solution2),
	})
}

func part1(monkeys []Monkey) int {
	const ROUNDS = 20
	inspections := make([]int, len(monkeys))

	for i := 0; i < ROUNDS; i++ {
		for j, monkey := range monkeys {
			for len(monkey.items) > 0 {
				// remove item from the current monkey's list
				item := monkey.shift()

				// Inspect and perform worry operation, divide outcome by 3
				updatedItem := monkey.operation(item) / 3

				// Test and throw to new monkey
				if updatedItem%monkey.divisionTestNumber == 0 {
					(monkeys)[monkey.positiveTestMonkey].addItem(updatedItem)
				} else {
					(monkeys)[monkey.negativeTestMonkey].addItem(updatedItem)
				}

				inspections[j] += 1
			}

			monkeys[j] = monkey
		}
	}

	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func part2(monkeys []Monkey) int {
	const ROUNDS = 10000
	inspections := make([]int, len(monkeys))

	// Based on Chinese Remainder Theorem
	moduliProduct := 1
	for _, monkey := range monkeys {
		moduliProduct *= monkey.divisionTestNumber
	}

	for i := 0; i < ROUNDS; i++ {
		for j, monkey := range monkeys {
			for len(monkey.items) > 0 {
				// remove item from the current monkey's list
				item := monkey.shift()

				// Inspect and perform worry operation
				updatedItem := monkey.operation(item) % moduliProduct

				// Test and throw to new monkey
				if updatedItem%monkey.divisionTestNumber == 0 {
					(monkeys)[monkey.positiveTestMonkey].addItem(updatedItem)
				} else {
					(monkeys)[monkey.negativeTestMonkey].addItem(updatedItem)
				}

				inspections[j] += 1
			}

			monkeys[j] = monkey
		}
	}

	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func (monkey *Monkey) shift() int {
	item := monkey.items[0]
	monkey.items = monkey.items[1:]

	return item
}

func (monkey *Monkey) addItem(item int) {
	monkey.items = append(monkey.items, item)
}
