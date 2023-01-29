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
	solution1 := part1(moves)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %v", solution1),
	})
}

func part1(moves []Move) int {
	tailVisitedMap := makeMoves(moves)

	return len(tailVisitedMap)
}

func makeMoves(moves []Move) TailMap {
	tailVisitedMap := TailMap{"0,0": void{}}
	headPosition := Coord{x: 0, y: 0}
	tailPosition := Coord{x: 0, y: 0}

	for _, move := range moves {
		fmt.Println("move:", move)

		for i := move.steps; i > 0; i-- {
			headPosition.moveHead(move.direction)

			fmt.Println("head position:", headPosition)

			if absDiff(tailPosition.x, headPosition.x) > 1 || absDiff(tailPosition.y, headPosition.y) > 1 {
				tailPosition.moveTail(headPosition, move.direction)

				tailVisited := fmt.Sprintf("%d,%d", tailPosition.x, tailPosition.y)
				tailVisitedMap[tailVisited] = void{}
			}

			fmt.Println("tail position", tailPosition)
		}
	}

	return tailVisitedMap
}

func (tailPosition *Coord) moveTail(headPosition Coord, direction string) {
	switch direction {
	case "U":
		tailPosition.x = headPosition.x
		tailPosition.y = tailPosition.y + 1
	case "D":
		tailPosition.x = headPosition.x
		tailPosition.y = tailPosition.y - 1
	case "R":
		tailPosition.x = tailPosition.x + 1
		tailPosition.y = headPosition.y
	case "L":
		tailPosition.x = tailPosition.x - 1
		tailPosition.y = headPosition.y
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
