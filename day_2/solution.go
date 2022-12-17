package main

import (
	"advent-of-code-2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// A : X -> rock
	// B : Y -> papers
	// C : Z -> scissors

	solutions := readInput()
	utils.PrintSolution(solutions)
}

func readInput() *[]string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Printf("unable to read file\n%s", err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	win := []string{"A Y", "B Z", "C X"}
	loss := []string{"A Z", "B X", "C Y"}
	draw := []string{"A X", "B Y", "C Z"}

	bonus := map[string]int{"X": 1, "Y": 2, "Z": 3}

	score := 0

	for scanner.Scan() {
		round := scanner.Text()
		roundscore := 0

		if contains(win, round) {
			roundscore = 6 + getBonusScore(round, bonus)
			score += roundscore
		}

		if contains(draw, round) {
			roundscore = 3 + getBonusScore(round, bonus)
			score += roundscore
		}

		if contains(loss, round) {
			roundscore = 0 + getBonusScore(round, bonus)
			score += roundscore
		}

		//fmt.Printf("round score: %d\n"+"total score %d\n", roundscore, score)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &[]string{fmt.Sprintf("score: %d", score)}
}

func getBonusScore(round string, bonus map[string]int) int {
	sign := strings.Split(round, " ")[1]
	return bonus[sign]
}

func contains(list []string, element string) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}

	return false
}
