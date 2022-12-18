package main

import (
	"advent-of-code-2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
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

	solution1 := part1()
	solution2 := part2()
	solutions := append(*solution1, *solution2...)

	utils.PrintSolution(&solutions)
}

func part1() *[]string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Printf("unable to read file\n%s", err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	// A : X -> rock
	// B : Y -> papers
	// C : Z -> scissors

	win := []string{"A Y", "B Z", "C X"}
	loss := []string{"A Z", "B X", "C Y"}
	draw := []string{"A X", "B Y", "C Z"}

	score := 0

	for scanner.Scan() {
		round := scanner.Text()
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &[]string{fmt.Sprintf("part 1 score: %d", score)}
}

func part2() *[]string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Printf("unable to read file\n%s", err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	// A -> rock
	// B -> papers
	// C -> scissors

	// X -> loss
	// Y -> draw
	// Z -> win

	signs := map[string]int{"A": rock, "B": paper, "C": scissors}
	outcomes := map[string]int{"X": loss, "Y": draw, "Z": win}

	score := 0

	for scanner.Scan() {
		round := scanner.Text()

		split := strings.Split(round, " ")
		oponentSign, outcome := split[0], split[1]

		ownSignScore := 0

		if outcomes[outcome] == loss {
			ownSignScore = getLosingSignScore(signs[oponentSign])
		}

		if outcomes[outcome] == draw {
			ownSignScore = signs[oponentSign]
		}

		if outcomes[outcome] == win {
			ownSignScore = getWinningSignScore(signs[oponentSign])
		}

		roundScore := outcomes[outcome] + ownSignScore
		score += roundScore

		//fmt.Printf("round score: %d\n"+"total score %d\n", roundScore, score)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
