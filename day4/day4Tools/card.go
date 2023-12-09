package day4Tools

import (
	"strconv"
	"strings"
)

type Card struct {
	Index         int
	winningNumber []int
	lotteryNumber []int
	value         int
}

func parseInputToCard(inputString string) Card {
	var id int = getIdFromString(inputString)
	var winningNumber []int
	var lotteryNumber []int
	winningNumber, lotteryNumber = parseNumbers(inputString)
	return Card{id, winningNumber, lotteryNumber, 1}
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
	var winningNumbers []int
	var lotteryNumber []int
	var numberString string = inputString[strings.IndexByte(inputString, ':')+1:]
	var numberStringArray []string = strings.Split(numberString, " | ")
	winningNumbers = parseNumberString(numberStringArray[0])
	lotteryNumber = parseNumberString(numberStringArray[1])
	return winningNumbers, lotteryNumber
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
