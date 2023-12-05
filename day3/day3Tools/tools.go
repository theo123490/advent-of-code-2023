package day3Tools

import (
	"fmt"
	"unicode"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Item struct {
	Name    string
	isDigit bool
	X       []int
	Y       int
}

func GetFinalResult(inputFile string) int {
	itemSlices := getSchematicCoordinates(inputFile)
	symbolSlice := getSymbol(itemSlices)
	fmt.Println(symbolSlice)
	for i := range symbolSlice {
		fmt.Println(symbolSlice[i].Name)
	}
	fmt.Println("*******")
	numberSlice := getNumbers(itemSlices)
	fmt.Println(numberSlice)
	for i := range numberSlice {
		fmt.Println(numberSlice[i].Name)
	}
	return 0
}

func getSchematicCoordinates(inputFile string) []Item {
	scanner, file := commonTools.FileReader(inputFile)
	defer file.Close()
	var itemSlice []Item = make([]Item, 0)
	var currentY = 0
	for scanner.Scan() {
		var inputString string = scanner.Text()

		var itemRunes []rune = make([]rune, 0)
		var isObject bool = false
		var item Item = Item{}

		for currentX, value := range inputString {
			if value == '.' {
				if isObject {
					item.Name = string(itemRunes)
					itemSlice = append(itemSlice, item)

					itemRunes = make([]rune, 0)
					isObject = false
					item = Item{}
				}
			} else {
				if unicode.IsDigit(value) != item.isDigit && len(itemRunes) > 0 {
					item.Name = string(itemRunes)
					itemSlice = append(itemSlice, item)

					itemRunes = make([]rune, 0)
					isObject = false
					item = Item{}
				}

				isObject = true
				item.X = append(item.X, currentX)
				item.Y = currentY
				item.isDigit = unicode.IsDigit(value)

				itemRunes = append(itemRunes, value)
			}
		}
		currentY += 1
	}
	fmt.Println(itemSlice)
	return itemSlice
}

func getSymbol(allItemSlice []Item) []*Item {
	var symbolSlices []*Item
	for i := range allItemSlice {
		if !(allItemSlice[i].isDigit) {
			symbolSlices = append(symbolSlices, &allItemSlice[i])
		}
	}
	return symbolSlices
}

func getNumbers(allItemSlice []Item) []*Item {
	var numberSlices []*Item
	for i := range allItemSlice {
		if allItemSlice[i].isDigit {
			numberSlices = append(numberSlices, &allItemSlice[i])
		}
	}
	return numberSlices
}
