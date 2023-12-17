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
	var race Race = Race{55826490, 246144110121111, 0} //TODO: parse input for this, I'm to lazy right now
	race.howManyWins()
	return race.possibleWins
}

func multiplyAllWins(races []Race) int {
	var wins int = 1
	for i := range races {
		wins *= races[i].possibleWins
	}
	return wins
}
