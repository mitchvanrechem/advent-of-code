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

	return strings.Split(string(data), "\n")
}

func ReadInputAsBytes(filePath string) [][]byte {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("unable to read file: ", err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	byteLines := [][]byte{}

	// scanner.Split() can be called with a SplitFunc to split up data:
	// scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)

	// By default scanner.Scan() will split the data similar to bufio.ScanLines

	// Any Custom split function can be written as long as it has the following signature:
	// func(data []byte, atEOF bool) (advance int, token []byte, err error)

	// While scanning the data can be output as:
	// a string with .Text()
	// a bytes slice with .Bytes()

	for scanner.Scan() {
		byteLines = append(byteLines, scanner.Bytes())
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
