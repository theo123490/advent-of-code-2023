package day6Tools

import (
	"strconv"
	"strings"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Race struct {
	time     int
	distance int
}

func getRaces(inputFile string) []Race {
	var races []Race
	var rawInputSlice []string = commonTools.CreateInputStringSlice(inputFile)
	var rawTime []string = strings.Fields(rawInputSlice[0])
	var rawDistance []string = strings.Fields(rawInputSlice[1])
	for i := 1; i <= len(rawTime)-1; i++ {
		time, err := strconv.Atoi(rawTime[i])
		commonTools.PrintErr(err)
		distance, err := strconv.Atoi(rawDistance[i])
		commonTools.PrintErr(err)
		var race Race = Race{time, distance}
		races = append(races, race)
	}
	return races
}
