package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInputAsStrings("input.txt")
	cards := parseInput(lines)

	solution1, solution2 := part1And2(cards)

	fmt.Printf("Solution to part 1: %v\n", solution1)
	fmt.Printf("Solution to part 2: %v\n", solution2)
}

type Card struct {
	winning []int
	owned   []int
}

func parseInput(lines []string) []Card {
	cards := make([]Card, len(lines))
	numberExp := regexp.MustCompile(`[0-9]+`)

	for i, line := range lines {
		cardContent := strings.Split(line, ":")[1]
		splitCardContent := strings.Split(cardContent, " | ")

		winningAsStrings := numberExp.FindAllString(splitCardContent[0], -1)
		ownedAsStrings := numberExp.FindAllString(splitCardContent[1], -1)

		card := Card{
			winning: make([]int, len(winningAsStrings)),
			owned:   make([]int, len(ownedAsStrings)),
		}

		for j := 0; j < len(winningAsStrings); j++ {
			card.winning[j], _ = strconv.Atoi(winningAsStrings[j])
		}

		for j := 0; j < len(ownedAsStrings); j++ {
			if ownedAsStrings[j] == " " {
				continue
			}
			card.owned[j], _ = strconv.Atoi(ownedAsStrings[j])
		}

		cards[i] = card
	}

	return cards
}

func part1And2(cards []Card) (int, int) {
	// part 1 vars
	totalPoints := 0

	// part 2 vars
	totalCards := 0
	copiesMap := make(map[int]int)

	for i, card := range cards {
		count := 0

		fmt.Println(card)

		for _, ownNum := range card.owned {
			if slices.Contains(card.winning, ownNum) {
				count++
			}
		}

		updatePart1TotalPoints(count, &totalPoints)
		updatePart2TotalCards(count, i+1, &copiesMap, &totalCards)
	}

	return totalPoints, totalCards
}

func updatePart1TotalPoints(count int, totalPoints *int) {
	points := 0

	if count == 1 {
		points += 1
	} else {
		points += int(math.Pow(2, float64(count-1)))
	}

	*totalPoints += points
}

func updatePart2TotalCards(count int, cardId int, copiesMap *map[int]int, totalCards *int) {
	copiesToAdd := 0

	if (*copiesMap)[cardId] != 0 {
		copiesToAdd = (*copiesMap)[cardId]
	}

	for j := 1; j <= count; j++ {
		(*copiesMap)[cardId+j] += (1 + copiesToAdd)
	}

	fmt.Println(copiesMap)

	*totalCards += (1 + copiesToAdd)
}
