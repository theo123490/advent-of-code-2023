package day6Tools

import "fmt"

func GetFinalResult(inputFile string) int {
	var races []Race = getRaces(inputFile)
	for i := range races {
		races[i].howManyWins()
	}
	fmt.Println(races)
	return multiplyAllWins(races)
}

func GetFinalResult2(inputFile string) int {
	return 0
}

func multiplyAllWins(races []Race) int {
	var wins int = 1
	for i := range races {
		wins *= races[i].possibleWins
	}
	return wins
}
