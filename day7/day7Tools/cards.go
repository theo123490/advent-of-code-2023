package day7Tools

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Card rune

type CardHand struct {
	cards        []Card
	cardValueMap map[Card]int
}

func CreateCardValueMap() map[Card]int {
	var cardValueMap map[Card]int = map[Card]int{
		Card("A"[0]): 12,
		Card("K"[0]): 11,
		Card("Q"[0]): 10,
		Card("J"[0]): 9,
		Card("T"[0]): 8,
		Card("9"[0]): 7,
		Card("8"[0]): 6,
		Card("7"[0]): 5,
		Card("6"[0]): 4,
		Card("5"[0]): 3,
		Card("4"[0]): 2,
		Card("3"[0]): 1,
		Card("2"[0]): 0,
	}
	return cardValueMap
}

func (ch CardHand) IsHigher(otherValue commonTools.Valuables) bool {
	var otherCh CardHand
	var ok bool

	if otherCh, ok = otherValue.(CardHand); !ok {
		fmt.Errorf("Can only compare with other cardHands not")
	}

	if &ch.cardValueMap != &otherCh.cardValueMap {
		fmt.Errorf("cardValueMap doesnt match")
	}
	var chMap map[Card]int = ch.cardValueMap

	return chMap[ch.cards[0]] > chMap[otherCh.cards[0]]
}
