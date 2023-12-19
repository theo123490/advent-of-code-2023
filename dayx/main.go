package main

import (
	"fmt"
	"strconv"

	"github.com/theo123490/advent-of-code-2023/dayx/dayxTools"
)

func main() {
	// var inputFile string = "dayx/example.txt"
	var inputFile string = "dayx/input.txt"
	var finalInt int = dayxTools.GetFinalResult(inputFile)
	var finalInt2 int = dayxTools.GetFinalResult2(inputFile)
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt))
	fmt.Println("$$$$$$$")
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt2))
	fmt.Println("$$$$$$$")
}
