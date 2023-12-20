package day7Tools

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type myInt int

func (i myInt) IsHigher(otherValue commonTools.Valuables) bool {
	ii := otherValue.(myInt)
	return i > ii
}

func GetFinalResult(inputFile string) int {
	var testSlice []commonTools.Valuables = []commonTools.Valuables{myInt(1), myInt(3), myInt(2), myInt(323), myInt(4), myInt(6), myInt(8), myInt(3), myInt(7), myInt(8), myInt(645), myInt(3643), myInt(4)}
	var result []commonTools.Valuables = commonTools.BubbleSort(testSlice)
	fmt.Println(result)
	return 0
}

func GetFinalResult2(inputFile string) int {
	return 0
}
