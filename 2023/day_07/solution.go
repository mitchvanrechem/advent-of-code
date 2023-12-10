package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

var cardStrengthMapP1 = map[rune]int{'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8, '9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0}
var cardStrengthMapP2 = map[rune]int{'A': 12, 'K': 11, 'Q': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1, 'J': 0}

func main() {
	lines := utils.ReadInputAsStrings("input.txt")
	hands := parseInput(lines)

	solution1 := part1and2(hands, false)
	solution2 := part1and2(hands, true)

	fmt.Printf("Solution to part 1: %v\n", solution1)
	fmt.Printf("Solution to part 2: %v\n", solution2)
}

func part1and2(hands []Hand, p2 bool) int {
	buckets := map[string][]Hand{
		"fiveOfAKind":  {},
		"fourOfAKind":  {},
		"fullHouse":    {},
		"threeOfAKind": {},
		"twoPair":      {},
		"onePair":      {},
		"highCard":     {},
	}

	for _, hand := range hands {
		handMap := make(map[rune]int, 5)

		highestDuplicate := 1
		secondHighestDuplicate := 0

		for _, card := range hand.cards {
			if _, ok := handMap[card]; !ok {
				handMap[card] = 1
			} else {
				handMap[card]++
			}

			if handMap[card] > highestDuplicate {
				highestDuplicate = handMap[card]
			} else if handMap[card] > secondHighestDuplicate {
				secondHighestDuplicate = handMap[card]
			}
		}

		if val, ok := handMap['J']; p2 && ok {
			// when 'J' is present it should always be in service of another card
			// only when J is the only kind of card, it can be treated as the highest duplicate

			if val == highestDuplicate {
				// if the 'J' duplicate amount is added to the second highest duplicate,
				// the second highest duplicate might become the highest, so that value should be set as the highest duplicate
				// for the rest of the algorithm to work
				change := secondHighestDuplicate + val

				if change > highestDuplicate {
					highestDuplicate = change
				}

			} else {
				highestDuplicate += val
			}

		}

		hand.determineBucket(highestDuplicate, secondHighestDuplicate, &buckets)
	}

	for key, val := range buckets {
		slices.SortFunc(val, func(a Hand, b Hand) int {
			for i := 0; i < len(a.cards); i++ {
				var aStr int
				var bStr int
				if p2 {
					aStr = cardStrengthMapP2[rune(a.cards[i])]
					bStr = cardStrengthMapP2[rune(b.cards[i])]
				} else {
					aStr = cardStrengthMapP1[rune(a.cards[i])]
					bStr = cardStrengthMapP1[rune(b.cards[i])]
				}

				if aStr < bStr {
					return -1
				} else if aStr > bStr {
					return 1
				}
			}

			return 0
		})

		slices.Reverse(val)
		buckets[key] = val
	}

	allHands := concatAllBuckets([][]Hand{
		buckets["fiveOfAKind"],
		buckets["fourOfAKind"],
		buckets["fullHouse"],
		buckets["threeOfAKind"],
		buckets["twoPair"],
		buckets["onePair"],
		buckets["highCard"],
	}, len(hands))

	var totalWinnigs int = 0

	for i := 0; i < len(allHands); i++ {
		totalWinnigs += allHands[i].bid * (len(allHands) - i)
	}

	return totalWinnigs
}

func parseInput(lines []string) []Hand {
	hands := make([]Hand, len(lines))

	for i, line := range lines {
		splitLine := strings.Split(line, " ")
		bid, _ := strconv.Atoi(splitLine[1])

		hands[i] = Hand{cards: splitLine[0], bid: bid}
	}

	return hands
}

func concatAllBuckets(buckets [][]Hand, len int) []Hand {
	concat := make([]Hand, len)

	var i int

	for _, s := range buckets {
		i += copy(concat[i:], s)
	}

	return concat
}

func (hand *Hand) determineBucket(highest int, secondHighest int, buckets *map[string][]Hand) {

	if highest >= 5 {
		(*buckets)["fiveOfAKind"] = append((*buckets)["fiveOfAKind"], *hand)
		return
	}

	if highest == 4 {
		(*buckets)["fourOfAKind"] = append((*buckets)["fourOfAKind"], *hand)
		return
	}

	if highest == 3 && secondHighest == 2 {
		(*buckets)["fullHouse"] = append((*buckets)["fullHouse"], *hand)
		return
	}

	if highest == 3 {
		(*buckets)["threeOfAKind"] = append((*buckets)["threeOfAKind"], *hand)
		return
	}

	if highest == 2 && secondHighest == 2 {
		(*buckets)["twoPair"] = append((*buckets)["twoPair"], *hand)
		return
	}

	if highest == 2 {
		(*buckets)["onePair"] = append((*buckets)["onePair"], *hand)
		return
	}

	(*buckets)["highCard"] = append((*buckets)["highCard"], *hand)
}
