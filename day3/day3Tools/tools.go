package day3Tools

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Item struct {
	Name         string
	isDigit      bool
	coordinate   Coordinate
	isPartNumber bool
}

type Coordinate struct {
	X []int
	Y []int
}

func GetFinalResult(inputFile string) int {
	var finalInt int = 0
	itemSlices := getSchematicCoordinates(inputFile)
	symbolSlice := getSymbol(itemSlices)
	numberSlice := getNumbers(itemSlices)
	fmt.Println(symbolSlice)
	for i := range symbolSlice {
		fmt.Println(symbolSlice[i].Name)
	}
	// fmt.Println("*******")
	// fmt.Println(numberSlice)
	// for i := range numberSlice {
	// 	fmt.Println(numberSlice[i].Name)
	// }

	var symbolSurroundingSlice []Coordinate
	for _, symbol := range symbolSlice {
		symbolX := symbol.coordinate.X[0]
		symbolY := symbol.coordinate.Y[0]
		symbolSurroundingSlice = append(symbolSurroundingSlice, getCoordinateSurrounding(symbolX, symbolY))
	}

	flagPartNumber(symbolSurroundingSlice, numberSlice)
	for i := range numberSlice {
		if numberSlice[i].isPartNumber {
			partNumberInt, err := strconv.Atoi(numberSlice[i].Name)
			// fmt.Println(partNumberInt)
			if err != nil {
				panic(err)
			}
			finalInt += partNumberInt
		}
	}

	return getGearRatioSum(itemSlices, numberSlice)
	// return finalInt
}

func getGearRatioSum(itemSlices []Item, numberSlice []*Item) int {
	gearSlice := getGears(itemSlices)
	fmt.Println(gearSlice)
	for i := range gearSlice {
		fmt.Println(gearSlice[i].Name)
	}

	var gearRatioSum int = 0
	for _, symbol := range gearSlice {
		var gearSurroundingSlice []Coordinate = []Coordinate{}
		// fmt.Println(symbol)
		var gearNumbers []int = []int{}
		gearX := symbol.coordinate.X[0]
		gearY := symbol.coordinate.Y[0]
		gearSurroundingSlice = append(gearSurroundingSlice, getCoordinateSurrounding(gearX, gearY))
		for _, surroundingCoordinate := range gearSurroundingSlice {
			for numberIndex, _ := range numberSlice {
				if numberSlice[numberIndex].isInCoordinate(surroundingCoordinate) {
					currentGearNumber, _ := strconv.Atoi(numberSlice[numberIndex].Name)
					gearNumbers = append(gearNumbers, currentGearNumber)
				}
			}
		}
		fmt.Println(gearNumbers)
		if len(gearNumbers) == 2 {
			gearRatio := gearNumbers[0] * gearNumbers[1]
			gearRatioSum += gearRatio
		} else {
			fmt.Println(gearNumbers)
		}

	}
	return gearRatioSum
}

func getGears(allItemSlice []Item) []*Item {
	var gearSlices []*Item
	for i := range allItemSlice {
		if allItemSlice[i].Name == "*" {
			gearSlices = append(gearSlices, &allItemSlice[i])
		}
	}
	return gearSlices
}

func flagPartNumber(symbolSurroundingSlice []Coordinate, numberSlice []*Item) {
	for _, surroundingCoordinate := range symbolSurroundingSlice {
		for numberIndex, numberPointer := range numberSlice {
			if numberSlice[numberIndex].isInCoordinate(surroundingCoordinate) {
				numberPointer.isPartNumber = true
			}
		}
	}
}

func (coordinate Coordinate) getRawCoordinate() [][]int {
	var itemRawCoordinates [][]int
	for _, x := range coordinate.X {
		for _, y := range coordinate.Y {
			itemRawCoordinates = append(itemRawCoordinates, []int{x, y})
		}
	}

	return itemRawCoordinates
}

func (item Item) isInCoordinate(coordinate Coordinate) bool {
	itemRawCoordinates := item.coordinate.getRawCoordinate()
	referenceRawCoordinate := coordinate.getRawCoordinate()
	for _, rawCoordinate := range referenceRawCoordinate {
		for _, itemCoordinate := range itemRawCoordinates {
			if rawCoordinate[0] == itemCoordinate[0] && rawCoordinate[1] == itemCoordinate[1] {
				return true
			}
		}
	}

	return false
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
				item.coordinate.X = append(item.coordinate.X, currentX)
				item.coordinate.Y = append(item.coordinate.Y, currentY)
				item.isDigit = unicode.IsDigit(value)

				itemRunes = append(itemRunes, value)
			}
		}

		if isObject {
			item.Name = string(itemRunes)
			itemSlice = append(itemSlice, item)

			isObject = false
			item = Item{}
		}

		currentY += 1
	}
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

func getCoordinateSurrounding(x int, y int) Coordinate {
	var coordinate Coordinate = Coordinate{}
	var surroundingTranslateMatrix []int = []int{-1, 0, 1}

	for _, xTranslate := range surroundingTranslateMatrix {
		coordinate.X = append(coordinate.X, x+xTranslate)
	}
	for _, yTranslate := range surroundingTranslateMatrix {
		coordinate.Y = append(coordinate.Y, y+yTranslate)
	}

	return coordinate
}
