package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInputAsStrings(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	inputLines := strings.Split(string(data), "\n")
	lines := []string{}

	for _, inputLine := range inputLines {
		if inputLine != "" {
			lines = append(lines, inputLine)
		}
	}

	return lines
}

func ReadInputAsBytes(filePath string) [][]byte {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("unable to read file: ", err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	byteLines := [][]byte{}

	for scanner.Scan() {
		line := scanner.Bytes()
		byteLines = append(byteLines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return byteLines
}

func PrintInput(filePath string) {
	input := ReadInputAsStrings(filePath)
	fmt.Print(input)
}
