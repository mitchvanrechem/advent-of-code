package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PrintSolution(solutions *[]string) {
	output := strings.Join(*solutions, "\n")

	fmt.Println(output)
	printToOutputFile(output)
}

func printToOutputFile(output string) {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.WriteString(output); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote the solutions to output.txt file")
}
