package day7Tools

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Card rune
type CardHandType string

const highCard CardHandType = CardHandType("high_card")
const onePair CardHandType = CardHandType("one_pair")
const twoPair CardHandType = CardHandType("two_pair")
const threeOak CardHandType = CardHandType("three_oak")
const fullHouse CardHandType = CardHandType("full_house")
const fourOak CardHandType = CardHandType("four_oak")
const fiveOak CardHandType = CardHandType("five_oak")

type CardHand struct {
	cards            []Card
	cardValueMap     *map[Card]int
	cardHandType     CardHandType
	cardHandTypeRule map[CardHandType]int
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

func CreateCardHandTypeRuleMap() map[CardHandType]int {
	var cardHandTypeRule map[CardHandType]int = map[CardHandType]int{
		highCard:  1,
		onePair:   2,
		twoPair:   3,
		threeOak:  4,
		fullHouse: 5,
		fourOak:   6,
		fiveOak:   7,
	}
	return cardHandTypeRule
}

func (ch CardHand) IsHigher(otherValue commonTools.Valuables) bool {
	var otherCh CardHand
	var ok bool
	if otherCh, ok = otherValue.(CardHand); !ok {
		panic(fmt.Errorf("can only compare with other cardHands not"))
	}

	//TODO: create a check to make sure card value map matches
	var chMap map[Card]int = *ch.cardValueMap

	return chMap[ch.cards[0]] > chMap[otherCh.cards[0]]
}

func newCardHand(cards string, cardValueMap *map[Card]int, cardHandTypeRule map[CardHandType]int) CardHand {
	var currentCards []Card = make([]Card, 5)

	if len(cards) != 5 {
		panic(fmt.Errorf("cards length is not 5 its %d, make it 5 or else", len(cards)))
	}
	for i := range cards {
		currentCards[i] = Card(cards[i])
	}

	var cardHand CardHand = CardHand{currentCards, cardValueMap, "", cardHandTypeRule}
	cardHand.getCardHandType()

	return cardHand
}

func (ch *CardHand) getCardHandType() {
	var cardMap map[Card]int = ch.createCardMap()
	var cardHandTypeMap map[CardHandType]int = map[CardHandType]int{
		onePair:  0,
		threeOak: 0,
		fourOak:  0,
		fiveOak:  0,
	}
	ch.cardHandType = highCard
	for _, cardValue := range cardMap {
		if cardValue == 2 {
			cardHandTypeMap[onePair] += 1
		}
		if cardValue == 3 {
			cardHandTypeMap[threeOak] += 1
		}
		if cardValue == 4 {
			cardHandTypeMap[fourOak] += 1
		}
		if cardValue == 5 {
			cardHandTypeMap[fiveOak] += 1
		}
	}

	if cardHandTypeMap[onePair] == 1 {
		ch.cardHandType = onePair
	}
	if cardHandTypeMap[onePair] == 2 {
		ch.cardHandType = twoPair
	}
	if cardHandTypeMap[threeOak] >= 1 {
		ch.cardHandType = threeOak
	}
	if cardHandTypeMap[onePair] >= 1 && cardHandTypeMap[threeOak] >= 1 {
		ch.cardHandType = fullHouse
	}
	if cardHandTypeMap[fourOak] >= 1 {
		ch.cardHandType = fourOak
	}
	if cardHandTypeMap[fiveOak] >= 1 {
		ch.cardHandType = fiveOak
	}
}

func (ch *CardHand) createCardMap() map[Card]int {
	var cardMap map[Card]int = make(map[Card]int)
	for _, card := range ch.cards {
		cardMap[card] += 1
	}

	return cardMap
}
