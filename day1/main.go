package main

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

func main() {
	var inputFile string = "day1/input.txt"
	scanner, file := commonTools.FileReader(inputFile)
	defer file.Close()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
