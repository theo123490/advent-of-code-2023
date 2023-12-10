package day5Tools

import (
	"strconv"
	"strings"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Almanac struct {
	seed        []int
	soil        map[int]int
	fertilizer  map[int]int
	water       map[int]int
	light       map[int]int
	temperature map[int]int
	humidity    map[int]int
	location    map[int]int
}

func CreateAlamanac(inputFile string) Almanac {
	var almanac Almanac

	var rawAlmanac []string = commonTools.CreateInputStringSlice(inputFile)
	almanac.getSeeds(rawAlmanac)
	return almanac
}

func (almanac *Almanac) getSeeds(rawAlmanac []string) {
	var seedString string = rawAlmanac[0][strings.IndexByte(rawAlmanac[0], ':')+1:]
	var seedStringArray []string = strings.Fields(seedString)
	for _, currentSeedString := range seedStringArray {
		seedInt, err := strconv.Atoi(currentSeedString)
		if err != nil {
			panic(err)
		}

		almanac.seed = append(almanac.seed, seedInt)
	}
}
