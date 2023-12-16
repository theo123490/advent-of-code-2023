package day5Tools

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/theo123490/advent-of-code-2023/commonTools"
)

type Almanac struct {
	seed                  []int
	seedToSoil            []AlmanacMap
	soilToFertilizer      []AlmanacMap
	fertilizerToWater     []AlmanacMap
	waterToLight          []AlmanacMap
	lightToTemperature    []AlmanacMap
	temperatureToHumidity []AlmanacMap
	humidityToLocation    []AlmanacMap
}

type AlmanacMap struct {
	mapSource      string
	mapDestination string
	source         int
	destination    int
	plotRange      int
}

func CreateAlamanac(inputFile string) Almanac {
	var almanac Almanac

	var rawAlmanac []string = commonTools.CreateInputStringSlice(inputFile)
	almanac.getSeeds(rawAlmanac)
	var allAlmanacMap []AlmanacMap = getAllAlmanacMap(rawAlmanac)
	almanac.seedToSoil = getSourceToDest("seed", "soil", &allAlmanacMap)
	almanac.soilToFertilizer = getSourceToDest("soil", "fertilizer", &allAlmanacMap)
	almanac.fertilizerToWater = getSourceToDest("fertilizer", "water", &allAlmanacMap)
	almanac.waterToLight = getSourceToDest("water", "light", &allAlmanacMap)
	almanac.lightToTemperature = getSourceToDest("light", "temperature", &allAlmanacMap)
	almanac.temperatureToHumidity = getSourceToDest("temperature", "humidity", &allAlmanacMap)
	almanac.humidityToLocation = getSourceToDest("humidity", "location", &allAlmanacMap)
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

func getSourceToDest(source string, destination string, allAlmanacMap *[]AlmanacMap) []AlmanacMap {
	var sourceToDestMap []AlmanacMap
	for _, almanacMap := range *allAlmanacMap {
		if almanacMap.mapSource == source && almanacMap.mapDestination == destination {
			sourceToDestMap = append(sourceToDestMap, almanacMap)
		}
	}
	return sourceToDestMap
}

func getAllAlmanacMap(rawAlmanac []string) []AlmanacMap {
	var isMapArea bool = false
	var source string = ""
	var destination string = ""
	var allAlmanacMap []AlmanacMap

	for _, almanacLine := range rawAlmanac {
		if almanacLine == "" {
			isMapArea = false
		}

		if isMapArea {
			var almanacMap AlmanacMap
			almanacMap.mapSource = source
			almanacMap.mapDestination = destination
			almanacMap.parseAlmanacMap(almanacLine)
			allAlmanacMap = append(allAlmanacMap, almanacMap)
		}

		if isAlmanacMapIndex(almanacLine) {
			isMapArea = true
			source, destination = parseAlmanacMapIndex(almanacLine)
		}
	}
	return allAlmanacMap
}

func isAlmanacMapIndex(alamanacLine string) bool {
	flag, err := regexp.MatchString("^.+-to-.+ map", alamanacLine)
	printErr(err)
	return flag
}

func parseAlmanacMapIndex(alamanacLine string) (string, string) {
	var toMap []string = strings.Split(alamanacLine, "-to-")
	var source string = toMap[0]
	var destination string = toMap[1][:strings.IndexByte(toMap[1], ' ')]
	return source, destination
}

func (a *AlmanacMap) parseAlmanacMap(almanacMapLine string) {
	var almanacIndexSlice []string = strings.Fields(almanacMapLine)
	var err error
	a.source, err = strconv.Atoi(almanacIndexSlice[1])
	printErr(err)
	a.destination, err = strconv.Atoi(almanacIndexSlice[0])
	printErr(err)
	a.plotRange, err = strconv.Atoi(almanacIndexSlice[2])
	printErr(err)
}

func printErr(err error) {
	if err != nil {
		panic(err)
	}
}
