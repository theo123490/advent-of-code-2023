package main

import (
	"fmt"
	"strconv"

	"github.com/theo123490/advent-of-code-2023/day4/day4Tools"
)

func main() {
	// var inputFile string = "day4/example.txt"
	var inputFile string = "day4/input.txt"
	var finalInt int = day4Tools.GetFinalResult(inputFile)
	var finalInt2 int = day4Tools.GetFinalResult2(inputFile)
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt))
	fmt.Println("$$$$$$$")
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt2))
	fmt.Println("$$$$$$$")
}
