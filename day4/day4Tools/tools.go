package day4Tools

import (
	"github.com/theo123490/advent-of-code-2023/commonTools"
)

func GetFinalResult(inputFile string) int {
	var cardList []Card = getCardList(inputFile)
	var totalValue int = 0

	for _, card := range cardList {
		totalValue += card.value
	}
	return totalValue
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
