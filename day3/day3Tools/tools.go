package day3Tools

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Item struct {
	Item string
	X    []int
	Y    int
}

func GetFinalResult(inputFile string) int {
	getSchematicCoordinates(inputFile)
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
					item.Item = string(itemRunes)
					itemSlice = append(itemSlice, item)

					itemRunes = make([]rune, 0)
					isObject = false
					item = Item{}
				}
			} else {
				isObject = true

				item.X = append(item.X, currentX)
				item.Y = currentY

				itemRunes = append(itemRunes, value)
			}
		}
		currentY += 1
	}
	fmt.Println(itemSlice)
	return itemSlice
}
