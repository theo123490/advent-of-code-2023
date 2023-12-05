package day3Tools

import (
	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Coordinate struct {
	X int
	Y int
}

func GetFinalResult(inputFile string) int {
	getSchematicCoordinates(inputFile)
	return 0
}

func getSchematicCoordinates(inputFile string) map[string]Coordinate {
	scanner, file := commonTools.FileReader(inputFile)
	defer file.Close()
	var currentY = 0
	for scanner.Scan() {
		var inputString string = scanner.Text()
		for currentX, value := range inputString {
			println("TODO")
		}
		currentY += 1
	}
	return make(map[string]Coordinate)
}
