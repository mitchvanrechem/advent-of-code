package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprint("Error reading file: ", err))
	}


	seeds, maps := parseInput(file)

	solution1 := part1(seeds, maps)

	fmt.Printf("Solution to part 1: %v\n", solution1)
}

type mapEntry struct {
	destStrt int
	srcStrt  int
	len      int
}

type Map []mapEntry

func parseInput(file []byte) ([]int, []Map) {
	paragraphs := strings.Split(string(file), "\n\n")
	numberExp := regexp.MustCompile("[0-9]+")

	seeds := []int{}
	seedsAsStrings := numberExp.FindAllString(paragraphs[0], -1)

	for _, seedAsString := range seedsAsStrings {
		seed, _ := strconv.Atoi(seedAsString)
		seeds = append(seeds, seed)
	}

	maps := []Map{}

	for i := 1; i < len(paragraphs); i++ {
		maps = append(maps, parseParagraphLines(paragraphs[i]))
	}

	return seeds, maps
}

func parseParagraphLines(para string) Map {
	lines := strings.Split(para, "\n")
	mapEntries := []mapEntry{}

	for i := 1; i < len(lines); i++ {
		ranges := strings.Split(lines[i], " ")

		destination, _ := strconv.Atoi(ranges[0])
		source, _ := strconv.Atoi(ranges[1])
		length, _ := strconv.Atoi(ranges[2])

		mapEntries = append(mapEntries, mapEntry{
			destStrt: destination,
			srcStrt:  source,
			len:      length,
		})
	}

	return mapEntries
}

func part1(seeds []int, maps []Map) int {
	locations := []int{}

	for _, seed := range seeds {
		lookupNumber := seed

		for _, m := range maps {

			for _, mapEntry := range m {
				rangeStrt, rangeEnd := mapEntry.srcStrt, mapEntry.srcStrt+mapEntry.len

				if rangeStrt <= lookupNumber && lookupNumber <= rangeEnd {
					lookupNumber = mapEntry.destStrt + (lookupNumber - rangeStrt)
					break;
				}
				
			}
		}

		locations = append(locations, lookupNumber)
	}

	return slices.Min(locations)
}
