package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInputAsStrings(filePath string) []string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read file")
		panic(err)
	}

	return strings.Split(string(file), "\n")
}
