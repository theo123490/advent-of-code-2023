package day7Tools

import (
	"fmt"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

func GetFinalResult(inputFile string) int {
	var testSlice []int = []int{1, 3, 2, 323, 4, 6, 8, 3, 7, 8, 645, 3643, 4}
	var result []int = commonTools.BubbleSort(testSlice)
	fmt.Println(result)
	return 0
}

func GetFinalResult2(inputFile string) int {
	return 0
}
