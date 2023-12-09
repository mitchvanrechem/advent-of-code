package main

import (
	"fmt"
)

type RaceTable [][2]int

var raceTableP1 = RaceTable{{53, 333}, {83, 1635}, {72, 1289}, {88, 1532}}
var raceTableP2 = RaceTable{{53837288, 333163512891532}}

func main() {
	solution1, solution2 := part1and2()

	fmt.Printf("Solution to part 1: %v\n", solution1)
	fmt.Printf("Solution to part 2: %v\n", solution2)
}

func part1and2() (int, int) {
	// part 1
	waysToWinP1 := 1

	for _, race := range raceTableP1 {
		wins := 0

		getWinsCount(race[0], race[1], &wins)

		waysToWinP1 *= wins
	}

	// part 2
	waysTowinP2 := 0

	getWinsCount(raceTableP2[0][0], raceTableP2[0][1], &waysTowinP2)

	return waysToWinP1, waysTowinP2
}

func getWinsCount(maxTime int, minDistance int, wins *int) {
	for i := 1; i < maxTime; i++ {
		speed := i

		distance := speed * (maxTime - i)
		if distance > minDistance {
			*wins++
		}
	}
}
