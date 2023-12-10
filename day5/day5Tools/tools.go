package day5Tools

func GetFinalResult(inputFile string) int {
	var almanac Almanac = CreateAlamanac(inputFile)
	return len(almanac.seed)
}

func GetFinalResult2(inputFile string) int {
	return 0
}
