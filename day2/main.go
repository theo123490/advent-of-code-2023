package main

import (
	"fmt"
	"strconv"

	"github.com/theo123490/advent-of-code-2023/commonTools"
	"github.com/theo123490/advent-of-code-2023/day2/day2Tools"
)

func main() {
	// var inputFile string = "day2/example.txt"
	var inputFile string = "day2/input.txt"
	scanner, file := commonTools.FileReader(inputFile)
	defer file.Close()

	var finalInt int = 0
	existingBagMap := day2Tools.ExistingBagMap()
	for scanner.Scan() {
		var inputString string = scanner.Text()
		game := day2Tools.ParseGame(inputString)
		bagMap := day2Tools.CalculateMinimumCubeBag(game)
		if day2Tools.CheckIfBagMapPossible(existingBagMap, bagMap) {
			finalInt += game.Id
		}
	}
	fmt.Println("$$$$$$$")
	fmt.Println(strconv.Itoa(finalInt))
	fmt.Println("$$$$$$$")
}
