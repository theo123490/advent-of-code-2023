package day4Tools

import (
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
	var cardMap map[int]Card = makeCardMap(cardList)
	var cardCopies []Card = make([]Card, 0)

	for _, card := range cardList {
		cardCopies = append(cardCopies, card.getCardCopies(cardMap)...)
	}

	var allCard []Card = cardList

	for len(cardCopies) > 0 {
		var currentCard Card = cardCopies[0]
		var currentCardCopies []Card = currentCard.getCardCopies(cardMap)
		allCard = append(allCard, currentCard)
		cardCopies = append(cardCopies, currentCardCopies...)
		cardCopies = cardCopies[1:]
	}
	return len(allCard)
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
