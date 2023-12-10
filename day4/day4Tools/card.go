package day4Tools

import (
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Index             int
	winningNumberList []int
	lotteryNumberList []int
	pointValue        int
	winningPower      int
}

func parseInputToCard(inputString string) Card {
	var id int = getIdFromString(inputString)
	var winningNumber []int
	var lotteryNumber []int
	winningNumber, lotteryNumber = parseNumbers(inputString)

	var card Card = Card{id, winningNumber, lotteryNumber, -1, 0}
	card.calculateValue()
	return card
}

func getIdFromString(inputString string) int {
	const cardLetterCount int = 4
	var cardIdString string = inputString[:strings.IndexByte(inputString, ':')]
	cardIdString = strings.ReplaceAll(cardIdString, " ", "")
	id, err := strconv.Atoi(cardIdString[cardLetterCount:])
	if err != nil {
		panic(err)
	}
	return id
}

func parseNumbers(inputString string) ([]int, []int) {
	var winningNumbersList []int
	var lotteryNumberList []int
	var numberString string = inputString[strings.IndexByte(inputString, ':')+1:]
	var numberStringArray []string = strings.Split(numberString, " | ")
	winningNumbersList = parseNumberString(numberStringArray[0])
	lotteryNumberList = parseNumberString(numberStringArray[1])
	return winningNumbersList, lotteryNumberList
}

func parseNumberString(numberString string) []int {
	var numberSlice []int

	numberString = strings.TrimSpace(numberString)
	numberStringArray := strings.Fields(numberString)
	for _, numberString := range numberStringArray {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}

		numberSlice = append(numberSlice, number)
	}

	return numberSlice
}

func (card *Card) isNumberInLottery(referenceNumber int) bool {
	for _, lotteryNumber := range card.lotteryNumberList {
		if lotteryNumber == referenceNumber {
			return true
		}
	}

	return false
}

func (card *Card) calculateValue() {
	var winningPower int = 0
	for _, winningNumber := range card.winningNumberList {
		if card.isNumberInLottery(winningNumber) {
			winningPower += 1
		}
	}

	card.winningPower = winningPower

	if winningPower-1 < 0 {
		card.pointValue = 0
	} else {
		card.pointValue = int(math.Pow(2, float64(winningPower-1)))
	}
}

func makeCardMap(cardList []Card) map[int]Card {
	var cardMap map[int]Card = make(map[int]Card)
	for _, card := range cardList {
		cardMap[card.Index] = card
	}

	return cardMap
}

func (card Card) getCardCopies(cardMap map[int]Card) []Card {
	var cardCopies []Card = make([]Card, 0)
	for i := 1; i < card.winningPower+1; i++ {
		cardCopies = append(cardCopies, cardMap[card.Index+i])
	}

	return cardCopies
}
