package day6Tools

import "fmt"

func GetFinalResult(inputFile string) int {
	var races []Race = getRaces(inputFile)
	boat := NewBoat()
	boat.timePressed = 1
	fmt.Println(boat.isWinRace(races[0]))

	return len(races)
}

func GetFinalResult2(inputFile string) int {
	return 0
}
