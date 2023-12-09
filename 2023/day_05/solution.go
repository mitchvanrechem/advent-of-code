package main

import (
	"advent-of-code-2023/utils"
	"advent-of-code-2023/utils/cli"
	"advent-of-code-2023/utils/logging"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var args = cli.GetArgs()
var logger = logging.NewLogger(args["loggerEnabled"].(bool))

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprint("Error reading file: ", err))
	}

	seeds, maps := parseInput(file)

	solution1 := part1(seeds, maps)
	solution2 := part2(seeds, maps)

	fmt.Printf("Solution to part 1: %v\n", solution1)
	fmt.Printf("Solution to part 2: %v\n", solution2)
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
					break
				}
			}
		}

		locations = append(locations, lookupNumber)
	}

	return slices.Min(locations)
}

func part2(seeds []int, maps []Map) int {
	type Tuple [2]int

	currMapLookupTuples := []Tuple{}

	for i := 0; i < len(seeds); i += 2 {
		lookupStrt := seeds[i]
		lookupRange := seeds[i+1]

		currMapLookupTuples = append(currMapLookupTuples, Tuple{lookupStrt, lookupRange})
	}

	locations := []int{}

	for _, m := range maps {
		nextMapLookupTuples := []Tuple{}

		logger.Log("----------------------------------------------------")

		for len(currMapLookupTuples) > 0 {
			logger.Log(fmt.Sprintf("curr map lookups: %v", currMapLookupTuples))
			logger.Log(fmt.Sprintf("next map lookups: %v", nextMapLookupTuples))

			currTuple := utils.Pop(&currMapLookupTuples)
			lookupStrt, lookupRange := currTuple[0], currTuple[1]
			lookupEnd := lookupStrt + lookupRange - 1

			logger.Log(fmt.Sprintf("current tuple: %v", currTuple))
			logger.Log(fmt.Sprintf("tuple start: %v tuple end: %v", lookupStrt, lookupEnd))

			for i, mapEntry := range m {

				logger.Log(fmt.Sprintf("index: %v current map entry: %v", mapEntry, i))

				srcStrt, srcEnd := mapEntry.srcStrt, mapEntry.srcStrt+mapEntry.len-1
				destStrt := mapEntry.destStrt

				// if lookups (original seeds) fall out of the range of the source
				// add to the current map lookup tuple -> they might need more mapping in the next entries

				// if the lookups (original seeds) fall within the range of the source completely
				// map the complete look up range and add to the next map lookup tuples -> they don't need more mapping in next entries

				// if the lookups (original seeds) fall within the range of the source partially
				// add the mapped part to the next map lookup tuples -> they don't need more mapping in the next entries
				// add the non mapped range back to the current map lookup tuples -> they might need more mapping in the next entries

				// at the last map entry the current lookups have all been mapped, either to a dest or to their own values
				// the current map lookup tuples are added to the next map lookup tuples

				// Not mapped to destination case
				// Numbers are added as they are

				if srcStrt <= lookupStrt && lookupStrt <= srcEnd {
					// Completely mapped to destination case
					// Numbers are all mapped to the destination and added to the next map tuples
					if lookupEnd <= srcEnd {

						nextMapLookupTuples = append(nextMapLookupTuples, Tuple{
							destStrt + (lookupStrt - srcStrt),
							lookupRange,
						})
						break
					}

					// Partially mapped to destination case
					// Numbers up to the src end are mapped to the destination and added to the next map tuples
					// Numbers greater than the src end are not mapped and added as they are to the curr map tuples
					if lookupEnd > srcEnd {
						nextMapLookupTuples = append(nextMapLookupTuples, Tuple{
							destStrt + (lookupStrt - srcStrt),
							srcEnd - lookupStrt + 1,
						})
						currMapLookupTuples = append(currMapLookupTuples, Tuple{
							srcEnd + 1,
							lookupEnd - srcEnd + 1,
						})
						break
					}
				}

				if lookupStrt < srcStrt {
					// Partially mapped to destination case
					// Numbers from the source start up to the lookup end are mapped to the destination and added to the tuples
					// Numbers lesser than the src start are not mapped amd added as they are to the tuples
					if lookupEnd >= srcStrt && lookupEnd <= srcEnd {
						nextMapLookupTuples = append(nextMapLookupTuples, Tuple{
							destStrt,
							lookupStrt + lookupRange - srcStrt,
						})
						currMapLookupTuples = append(currMapLookupTuples, Tuple{
							lookupStrt,
							srcStrt - lookupStrt,
						})
						break
					}

					// Partially mapped to destination case
					// All src numbers are mapped to the destination and added to the tuples
					// lookup start is lesser than src start, the lookup numbers up to source start are not mapped and added as they are to the tuples
					// lookup end is greater than src end, the lookup numbers starting from the source end are not mapped and added as they are to the tuples
					if lookupEnd > srcEnd {
						nextMapLookupTuples = append(nextMapLookupTuples, Tuple{destStrt, mapEntry.len})
						currMapLookupTuples = append(currMapLookupTuples, Tuple{srcEnd, lookupEnd - srcEnd})
						currMapLookupTuples = append(currMapLookupTuples, Tuple{lookupStrt, srcStrt - lookupStrt})
						break
					}

				}

				// If last iteration of the map entries didn't break, the current tuple has no mapping and can be added to the next map tuples pool
				if i == len(m)-1 {
					if lookupStrt < srcStrt && lookupEnd < srcStrt {
						nextMapLookupTuples = append(nextMapLookupTuples, Tuple{lookupStrt, lookupRange})
						continue
					}
					if lookupStrt > srcEnd && lookupEnd > srcEnd {
						nextMapLookupTuples = append(nextMapLookupTuples, Tuple{lookupStrt, lookupRange})
						continue
					}
				}

			}

		}

		logger.Log(fmt.Sprintf("curr map lookups pre concat: %v", currMapLookupTuples))
		currMapLookupTuples = append(currMapLookupTuples, nextMapLookupTuples...)
		logger.Log(fmt.Sprintf("curr map lookups post concat: %v", currMapLookupTuples))
		logger.Log(fmt.Sprintf("next map lookups: %v", nextMapLookupTuples))
	}

	for _, tuple := range currMapLookupTuples {
		locations = append(locations, tuple[0])
	}

	return slices.Min(locations)
}
