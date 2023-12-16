package day5Tools

import "fmt"

func GetFinalResult(inputFile string) int {
	var almanac Almanac = CreateAlamanac(inputFile)
	var seedDestination []int = almanac.getAllSeedDestination()
	fmt.Println(seedDestination)
	return len(almanac.seed)
}

func GetFinalResult2(inputFile string) int {
	return 0
}
