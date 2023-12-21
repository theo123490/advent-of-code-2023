package day7Tools

import "fmt"

type myInt int

func GetFinalResult(inputFile string) int {
	var cardValueMap map[Card]int = CreateCardValueMap()
	var cardHandTypeRuleMap = CreateCardHandTypeRuleMap()
	var cardHandA CardHand = newCardHand("5A222", &cardValueMap, cardHandTypeRuleMap)
	var cardHandB CardHand = newCardHand("12345", &cardValueMap, cardHandTypeRuleMap)
	var cardHandC CardHand = newCardHand("22444", &cardValueMap, cardHandTypeRuleMap)
	fmt.Println(cardHandB.IsHigher(cardHandA))

	fmt.Println(cardHandA)
	fmt.Println(cardHandB)
	fmt.Println(cardHandC)
	return 0
}

func GetFinalResult2(inputFile string) int {
	return 0
}
