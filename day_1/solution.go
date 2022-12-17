package main

import (
	"advent-of-code-2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, advent of code 2022!")
	solutions := readInventory()

	utils.PrintSolution(solutions)
}

func readInventory() *[]string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Printf("unable to read file: %s", err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	calories_max := 0
	current_calories := 0
	elves := make([]int, 0)

	// scanner.Split() can be called with a SplitFunc to split up data:
	// scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)

	// By default scanner.Scan() will split the data similar to bufio.ScanLines

	// Any Custom split function can be written as long as it has the following signature:
	// func(data []byte, atEOF bool) (advance int, token []byte, err error)

	for scanner.Scan() {

		// While scanning the data can be output as: 
		// a string with .Text()
		// a bytes slice with .Bytes()

		calories_entry := scanner.Text()

		if calories_entry == "" {
			elves = append(elves, current_calories)
			//fmt.Printf("elf %d, with %d calories\n", len(elves), current_calories)

			if current_calories > calories_max {
				calories_max = current_calories
			}

			current_calories = 0
		} else {
			calories, _ := strconv.Atoi(calories_entry)
			current_calories += calories
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elves)

	return &[]string{
		fmt.Sprintf("Largest amount of calories: %d", calories_max),
		fmt.Sprintf("Top 3 elves: %d + %d + %d", elves[len(elves)-1], elves[len(elves)-2], elves[len(elves)-3]),
		fmt.Sprintf("Top 3 elves total amount: %d", elves[len(elves)-1]+elves[len(elves)-2]+elves[len(elves)-3]),
	}
}
