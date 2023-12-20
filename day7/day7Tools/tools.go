package day7Tools

import "fmt"

type myInt int

func GetFinalResult(inputFile string) int {
	var cardValueMap map[Card]int = CreateCardValueMap()
	var cardHandA CardHand = CardHand{
		[]Card{
			Card("2"[0]),
			Card("2"[0]),
			Card("4"[0]),
			Card("2"[0]),
			Card("2"[0]),
		}, cardValueMap}

	var cardHandB CardHand = CardHand{
		[]Card{
			Card("4"[0]),
			Card("2"[0]),
			Card("4"[0]),
			Card("2"[0]),
			Card("2"[0]),
		}, cardValueMap}

	fmt.Println(cardHandB.IsHigher(cardHandA))
	return 0
}

func GetFinalResult2(inputFile string) int {
	return 0
}
