package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Move struct {
	direction string
	steps     int
}

type Coord struct {
	x int
	y int
}

type void struct{}
type TailMap map[string]void

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")
	moves := parseInput(inputLines)

	// Kept in the logs for visualization (write to a file)
	solution1 := part1and2(moves, 2)
	solution2 := part1and2(moves, 10)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %v", solution1),
		fmt.Sprintf("part2: %v", solution2),
	})
}

func part1and2(moves []Move, ropeLength int) int {
	tailVisitedMap := makeMoves(moves, ropeLength)

	return len(tailVisitedMap)
}

func makeMoves(moves []Move, ropeLength int) TailMap {
	tailVisitedMap := TailMap{"0,0": void{}}
	rope := make([]Coord, ropeLength)

	for i := 0; i < ropeLength; i++ {
		rope[i] = Coord{x: 0, y: 0}
	}

	for _, move := range moves {
		fmt.Println("move:", move)

		for i := move.steps; i > 0; i-- {
			rope[0].moveHead(move.direction)

			fmt.Println("head position:", rope[0])

			for i := 1; i < ropeLength; i++ {
				if absDiff(rope[i-1].x, rope[i].x) > 1 || absDiff(rope[i-1].y, rope[i].y) > 1 {
					rope[i].moveTail(rope[i-1], move.direction)
				}
			}

			tailVisited := fmt.Sprintf("%d,%d", rope[ropeLength-1].x, rope[ropeLength-1].y)
			tailVisitedMap[tailVisited] = void{}

			fmt.Println("tail position", rope[ropeLength-1])
		}
	}

	return tailVisitedMap
}

func (curr *Coord) moveTail(prev Coord, direction string) {
	moveOverAxis(&curr.x, prev.x)
	moveOverAxis(&curr.y, prev.y)
}

func moveOverAxis(curr *int, prev int) {
	if absDiff(*curr, prev) >= 1 {
		if *curr > prev {
			*curr -= 1
		} else {
			*curr += 1
		}
	}
}

func absDiff(number1 int, number2 int) int {
	if number1 > number2 {
		return number1 - number2
	} else {
		return number2 - number1
	}
}

func (headPosition *Coord) moveHead(direction string) {
	switch direction {
	case "U":
		headPosition.y += 1
	case "D":
		headPosition.y -= 1
	case "R":
		headPosition.x += 1
	case "L":
		headPosition.x -= 1
	}
}

func parseInput(lines []string) []Move {
	moves := make([]Move, 0, len(lines))

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		steps, _ := strconv.Atoi(splitLine[1])

		move := Move{direction: splitLine[0], steps: steps}
		moves = append(moves, move)
	}

	return moves
}
