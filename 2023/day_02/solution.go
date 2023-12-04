package main

import (
	"advent-of-code-2023/utils"
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1(lines)
	solution2 := part2(lines)

	fmt.Printf("Solution to part 1: %v\n", solution1)
	fmt.Printf("Solution to part 2: %v\n", solution2)
}

type Game struct {
	id    int
	turns string
}

func part1(lines []string) int {
	allColoursExpression := regexp.MustCompile("[0-9]*")

	// Paranthese are used to group subexpressions, important for the later use
	// of "FindAllStringSubmatch", where the submatches are returned alongside
	// the complete match itself.
	redExpression := regexp.MustCompile("([0-9]*) red")
	greenExpression := regexp.MustCompile("([0-9]*) green")
	blueExpression := regexp.MustCompile("([0-9]*) blue")

	idsSum := 0

	for i, line := range lines {
		game := Game{id: i + 1, turns: strings.Split(line, ":")[1]}

		shownDice := allColoursExpression.FindAllString(game.turns, -1)

		maxAsString := slices.MaxFunc(shownDice, func(a, b string) int {
			aInt, _ := strconv.Atoi(a)
			bInt, _ := strconv.Atoi(b)

			return cmp.Compare(aInt, bInt)
		})

		maxDice, err := strconv.Atoi(maxAsString)

		if err == nil && maxDice > 14 {
			continue
		}

		// Max amount of each coloured cubes each turn in a game that can be shown:
		// 12 red cubes, 13 green cubes, 14 blue
		isPossibleForRed := game.isGamePossibleForGivenColour(redExpression, 12)
		isPossibleForGreen := game.isGamePossibleForGivenColour(greenExpression, 13)
		isPossibleForBlue := game.isGamePossibleForGivenColour(blueExpression, 14)

		if !<-isPossibleForRed || !<-isPossibleForGreen || !<-isPossibleForBlue {
			continue
		}

		idsSum += game.id
	}

	return idsSum
}

func (game *Game) isGamePossibleForGivenColour(colourExp *regexp.Regexp, maxPossibleDice int) <-chan bool {
	c := make(chan bool)

	go func() {
		colourSubmatches := colourExp.FindAllStringSubmatch(game.turns, -1)

		for _, subMatch := range colourSubmatches {
			diceValue, err := strconv.Atoi(subMatch[1])

			if err == nil && diceValue > maxPossibleDice {
				c <- false
			}
		}

		c <- true
	}()

	return c
}

func part2(lines []string) int {
	// Paranthese are used to group subexpressions, important for the later use
	// of "FindAllStringSubmatch", where the submatches are returned alongside
	// the complete match itself.
	redExpression := regexp.MustCompile("([0-9]*) red")
	greenExpression := regexp.MustCompile("([0-9]*) green")
	blueExpression := regexp.MustCompile("([0-9]*) blue")

	powerSum := 0

	for i, line := range lines {
		game := Game{id: i + 1, turns: strings.Split(line, ":")[1]}

		// Max amount of each coloured cubes each turn in a game that can be shown:
		// 12 red cubes, 13 green cubes, 14 blue
		minimumValueRed := game.getMinimumDiceValueForGivenColour(redExpression)
		minimumValueGreen := game.getMinimumDiceValueForGivenColour(greenExpression)
		minimumValueBlue := game.getMinimumDiceValueForGivenColour(blueExpression)

		power := <-minimumValueRed * <-minimumValueGreen * <-minimumValueBlue
		powerSum += power
	}

	return powerSum
}

func (game *Game) getMinimumDiceValueForGivenColour(colourExp *regexp.Regexp) <-chan int {

	c := make(chan int)

	go func() {
		colourSubmatches := colourExp.FindAllStringSubmatch(game.turns, -1)

		highestDiceValue, err := strconv.Atoi(colourSubmatches[0][1])
		if err != nil {
			panic(err)
		}

		for i := 1; i < len(colourSubmatches); i++ {
			diceValue, err := strconv.Atoi(colourSubmatches[i][1])

			if err == nil && diceValue > highestDiceValue {
				highestDiceValue = diceValue
			}
		}

		c <- highestDiceValue
	}()

	return c
}
