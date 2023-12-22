package day7Tools

func GetFinalResult(inputFile string) int {
	var cardHandSlice []CardHand = parseCardHands(inputFile)
	cardHandSlice = CardHandBubbleSort(cardHandSlice)

	// for _, card := range cardHandSlice {
	// 	fmt.Println(card.cards)
	// }
	return calculateBids(cardHandSlice)
}

func GetFinalResult2(inputFile string) int {
	return 0
}
