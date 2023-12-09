package main

import (
	"fmt"
	"strconv"

	"github.com/theo123490/advent-of-code-2023/day3/day3Tools"
)

func main() {
	// var inputFile string = "day3/example.txt"
	var inputFile string = "day3/input.txt"
	var finalInt int = day3Tools.GetFinalResult(inputFile)
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt))
	fmt.Println("$$$$$$$")
}
