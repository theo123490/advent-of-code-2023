package day6Tools

func GetFinalResult(inputFile string) int {
	var races []Race = getRaces(inputFile)
	return len(races)
}

func GetFinalResult2(inputFile string) int {
	return 0
}
