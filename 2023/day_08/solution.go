package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lines := utils.ReadInputAsStrings("input.txt")
	directions, network := parseInput(lines)

	solution1 := part1(directions, network)
	// solution2 := part2(directions, network)

	fmt.Printf("Solution to part 1: %v\n", solution1)
	// fmt.Printf("Solution to part 2: %v\n", solution2)
}

type Node string
type Network map[Node][2]Node

func part1(directions string, network Network) int {
	destReached := false
	steps := 0

	fmt.Println(directions)
	currNode := Node("AAA")

	for !destReached {
		for _, direction := range directions {
			if direction == 'L' {
				currNode = network[currNode][0]
			} else if direction == 'R' {
				currNode = network[currNode][1]
			}

			steps++

			if currNode == Node("ZZZ") {
				destReached = true
				break
			}

		}
	}

	return steps
}

func parseInput(lines []string) (string, Network) {
	directions := lines[0]
	lines = lines[2:]

	network := make(Network, len(lines))

	for _, line := range lines {
		splitLine := strings.Split(line, " = ")
		node := Node(splitLine[0])

		branches := strings.Split(splitLine[1], ", ")

		leftNode := Node(strings.TrimLeft(branches[0], "("))
		rightNode := Node(strings.TrimRight(branches[1], ")"))

		network[node] = [2]Node{leftNode, rightNode}
	}

	return directions, network
}
