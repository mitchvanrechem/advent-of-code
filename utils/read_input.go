package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

func PrintInput(filePath string) {
	input := ReadInput(filePath)
	fmt.Print(input)
}
