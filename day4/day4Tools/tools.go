package day4Tools

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

func GetFinalResult(inputFile string) int {
	var cardList []Card = getCardList(inputFile)
	var totalValue int = 0

	for _, card := range cardList {
		totalValue += card.pointValue
	}
	return totalValue
}

func GetFinalResult2(inputFile string) int {
	var cardList []Card = getCardList(inputFile)
	var totalValue int = 0
	var cardMap map[int]Card = makeCardMap(cardList)

	for _, card := range cardList {
		cardCopies := card.getCardCopies(cardMap)
		fmt.Println(cardCopies)
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
