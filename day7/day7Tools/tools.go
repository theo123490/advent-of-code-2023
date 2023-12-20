package day7Tools

import "fmt"

type myInt int

func GetFinalResult(inputFile string) int {
	var cardValueMap map[Card]int = CreateCardValueMap()
	var cardHandA CardHand = newCardHand("5A222", &cardValueMap, make(map[CardHandType]int))
	var cardHandB CardHand = newCardHand("42422", &cardValueMap, make(map[CardHandType]int))

	fmt.Println(cardHandB.IsHigher(cardHandA))
	return 0
}

func GetFinalResult2(inputFile string) int {
	return 0
}
