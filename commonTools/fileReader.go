package commonTools

import (
	"bufio"
	"log"
	"os"
)

func FileReader(inputFile string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scanner, file
}
