package main

import (
	"fmt"
	"strconv"

	"github.com/theo123490/advent-of-code-2023/day5/day5Tools"
)

func main() {
	var inputFile string = "day5/example.txt"
	// var inputFile string = "day5/input.txt"
	var finalInt int = day5Tools.GetFinalResult(inputFile)
	var finalInt2 int = day5Tools.GetFinalResult2(inputFile)
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt))
	fmt.Println("$$$$$$$")
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt2))
	fmt.Println("$$$$$$$")
}
