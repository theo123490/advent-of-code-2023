package main

import (
	"fmt"
	"strconv"

	"github.com/theo123490/advent-of-code-2023/commonTools"
	"github.com/theo123490/advent-of-code-2023/day1/day1Tools"
)

func main() {
	var inputFile string = "day1/input.txt"
	scanner, file := commonTools.FileReader(inputFile)
	defer file.Close()

	var finalInt int = 0
	var lineInt int
	for scanner.Scan() {
		fmt.Println("***********")
		var inputString string = scanner.Text()
		lineInt = day1Tools.StringParser(inputString)
		fmt.Println(strconv.Itoa(lineInt))
		finalInt += lineInt
	}
	fmt.Println(strconv.Itoa(finalInt))
}
