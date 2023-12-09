package day4Tools

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

func GetFinalResult(inputFile string) int {
	var cardSlice []Card = getCardList(inputFile)
	fmt.Println(cardSlice)
	return 0
}

func getCardList(inputFile string) []Card {
	var cardList []Card

	scanner, file := commonTools.FileReader(inputFile)
	defer file.Close()
	for scanner.Scan() {
		var inputString string = scanner.Text()
		card := parseInputToCard(inputString)
		cardList = append(cardList, card)
	}
	return cardList
}
