package main

import (
	"fmt"
	"strconv"

	"github.com/theo123490/advent-of-code-2023/day6/day6Tools"
)

func main() {
	var inputFile string = "day6/example.txt"
	// var inputFile string = "day6/input.txt"
	var finalInt int = day6Tools.GetFinalResult(inputFile)
	var finalInt2 int = day6Tools.GetFinalResult2(inputFile)
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt))
	fmt.Println("$$$$$$$")
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt2))
	fmt.Println("$$$$$$$")
}
