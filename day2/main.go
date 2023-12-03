package main

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
	"github.com/theo123490/advent-of-code-2023/day2/day2Tools"
)

func main() {
	var inputFile string = "day2/input.txt"
	scanner, file := commonTools.FileReader(inputFile)
	defer file.Close()

	// var finalInt int = 0
	// var lineInt int = 0
	for scanner.Scan() {
		var inputString string = scanner.Text()
		fmt.Println("***********")
		day2Tools.ParseGame(inputString)
	}
	// fmt.Println(strconv.Itoa(finalInt))
}
