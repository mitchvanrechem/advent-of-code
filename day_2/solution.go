package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

const (
	rock     int = 1
	paper    int = 2
	scissors int = 3
	loss     int = 0
	draw     int = 3
	win      int = 6
)

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1(inputLines)
	solution2 := part2(inputLines)
	solutions := append(*solution1, *solution2...)

	utils.PrintSolution(&solutions)
}

func part1(rounds []string) *[]string {

	// A : X -> rock
	// B : Y -> papers
	// C : Z -> scissors

	win := []string{"A Y", "B Z", "C X"}
	loss := []string{"A Z", "B X", "C Y"}
	draw := []string{"A X", "B Y", "C Z"}

	score := 0

	for _, round := range rounds {
		roundscore := 0

		if utils.Contains(win, round) {
			roundscore = 6 + getBonusScore(round)
			score += roundscore
		}

		if utils.Contains(draw, round) {
			roundscore = 3 + getBonusScore(round)
			score += roundscore
		}

		if utils.Contains(loss, round) {
			roundscore = 0 + getBonusScore(round)
			score += roundscore
		}

		//fmt.Printf("round score: %d\n"+"total score %d\n", roundscore, score)
	}

	return &[]string{fmt.Sprintf("part 1 score: %d", score)}
}

func part2(rounds []string) *[]string {
	// A -> rock
	// B -> papers
	// C -> scissors

	// X -> loss
	// Y -> draw
	// Z -> win

	signs := map[string]int{"A": rock, "B": paper, "C": scissors}
	outcomes := map[string]int{"X": loss, "Y": draw, "Z": win}

	score := 0

	for _, round := range rounds {
		split := strings.Split(round, " ")
		opponentSign, outcome := split[0], split[1]

		ownSignScore := 0

		if outcomes[outcome] == loss {
			ownSignScore = getLosingSignScore(signs[opponentSign])
		}

		if outcomes[outcome] == draw {
			ownSignScore = signs[opponentSign]
		}

		if outcomes[outcome] == win {
			ownSignScore = getWinningSignScore(signs[opponentSign])
		}

		roundScore := outcomes[outcome] + ownSignScore
		score += roundScore

		//fmt.Printf("round score: %d\n"+"total score %d\n", roundScore, score)
	}

	return &[]string{fmt.Sprintf("part 2 score: %d", score)}
}

func getBonusScore(round string) int {
	bonus := map[string]int{"X": rock, "Y": paper, "Z": scissors}

	sign := strings.Split(round, " ")[1]
	return bonus[sign]
}

func getWinningSignScore(openentSign int) int {
	if openentSign == rock {
		return paper
	}

	if openentSign == paper {
		return scissors
	}

	return rock
}

func getLosingSignScore(openentSign int) int {
	if openentSign == rock {
		return scissors
	}

	if openentSign == paper {
		return rock
	}

	return paper
}
