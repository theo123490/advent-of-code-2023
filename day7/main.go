package main

import (
	"fmt"
	"strconv"

	"github.com/theo123490/advent-of-code-2023/day7/day7Tools"
)

func main() {
	// var inputFile string = "day7/example.txt"
	var inputFile string = "day7/input.txt"
	var finalInt int = day7Tools.GetFinalResult(inputFile)
	var finalInt2 int = day7Tools.GetFinalResult2(inputFile)
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt))
	fmt.Println("$$$$$$$")
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt2))
	fmt.Println("$$$$$$$")
}
