package utils

import (
	"fmt"
	"os"
	"strings"
)

// Completely unnecessary feature, but wanted to try out writing to a file in GO
func PrintSolution(solutions *[]string) {
	output := strings.Join(*solutions, "\n")

	fmt.Println(output)
	printToOutputFile(output)
}

func printToOutputFile(output string) {
	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.WriteString(output); err != nil {
		panic(err)
	}

	fmt.Println("Wrote the solutions to output.txt file")
}
