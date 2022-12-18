package utils

import (
	"fmt"
	"log"
	"os"
)

func ReadInput(output string) {
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
