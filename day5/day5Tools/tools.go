package day5Tools

import (
	"github.com/theo123490/advent-of-code-2023/commonTools"
)

func GetFinalResult(inputFile string) int {
	var almanac Almanac = CreateAlamanac(inputFile)
	var seedDestination []int = almanac.getAllSeedDestination()
	return commonTools.GetMinimum(seedDestination)
}

func GetFinalResult2(inputFile string) int {
	var almanac Almanac = CreateAlamanac(inputFile)
	//This took too long, We should be able to find the solution by working backwards
	return almanac.getAllSeedRangeDestinationMin()
}
