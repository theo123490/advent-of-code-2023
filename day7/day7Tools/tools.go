package day7Tools

import "fmt"

func GetFinalResult(inputFile string) int {
	var cardHandSlice []CardHand = parseCardHands(inputFile)
	fmt.Println(cardHandSlice)
	return len(cardHandSlice)
}

func GetFinalResult2(inputFile string) int {
	return 0
}
