package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"regexp"
	"strconv"
)

const columns = 9

type Crate string
type Stack []Crate
type CargoDeck [columns]Stack

type Move struct {
	amount      int
	source      int
	destination int
}

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1(inputLines)
	solution2 := part2(inputLines)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %s", solution1),
		fmt.Sprintf("part2: %s", solution2),
	})
}

func part1(inputLines []string) string {
	deck, moves := parseInput(inputLines)

	for _, move := range moves {
		sourceStack := deck[move.source]
		destinationStack := deck[move.destination]

		for i := 0; i < move.amount; i++ {
			topCrate := removeTopCrates(&sourceStack, 1, false)
			destinationStack = append(destinationStack, topCrate...)
		}

		deck[move.source] = sourceStack
		deck[move.destination] = destinationStack
	}

	message := ""

	for _, stack := range deck {
		message += string(stack[len(stack)-1])
	}

	return message
}

func part2(inputLines []string) string {
	deck, moves := parseInput(inputLines)

	for _, move := range moves {
		sourceStack := deck[move.source]
		destinationStack := deck[move.destination]

		topCrates := removeTopCrates(&sourceStack, move.amount, true)
		destinationStack = append(destinationStack, topCrates...)

		deck[move.source] = sourceStack
		deck[move.destination] = destinationStack
	}

	message := ""

	for _, stack := range deck {
		message += string(stack[len(stack)-1])
	}

	return message
}

func removeTopCrates(stack *Stack, crateAmount int, part2 bool) Stack {
	stackHeight := len(*stack) - crateAmount
	topCrates := (*stack)[stackHeight:]
	*stack = (*stack)[:stackHeight]

	return topCrates
}

func parseMoves(lines []string) []Move {
	moves := []Move{}

	for _, line := range lines {
		re := regexp.MustCompile("[0-9]+")
		moveNumbers := re.FindAllString(line, -1)

		amount, _ := strconv.Atoi(moveNumbers[0])
		source, _ := strconv.Atoi(moveNumbers[1])
		destination, _ := strconv.Atoi(moveNumbers[2])

		// source - 1 and destination - 1, because stacks are 1 index based
		// but our arrays are 0 index based
		move := Move{
			amount:      amount,
			source:      source - 1,
			destination: destination - 1,
		}

		moves = append(moves, move)
	}

	return moves
}

func parseCrates(lines []string) CargoDeck {
	deck := CargoDeck{}
	rows := len(lines) - 1

	// Iterate over input line representing crate stack level from bottom to top
	for i := rows; i >= 0; i-- {
		setCratesPerRow(lines[i], &deck)
	}

	return deck
}

func setCratesPerRow(row string, deck *CargoDeck) {
	for i := 0; i < columns; i++ {
		// Representation of a Crate in a column is 4 characters wide
		rowPosition := i * 4

		if row[rowPosition] == '[' {
			crates := append(deck[i], Crate(row[rowPosition+1]))
			deck[i] = crates
		}
	}
}

func parseInput(lines []string) (CargoDeck, []Move) {
	deck := CargoDeck{}
	moves := []Move{}

	for i, line := range lines {
		if line == "" {
			// i-1 , cutting of the numbers underneath the crate schema
			deck = parseCrates(lines[:i-1])
			moves = parseMoves(lines[i+1:])

			break
		}
	}

	return deck, moves
}
