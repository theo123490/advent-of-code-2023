package day1Tools

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func StringParser(inputString string) int {
	var numbers string = findAllNumber(inputString)
	return parseNumberToInt(numbers)
}

func findAllNumber(inputString string) string {
	var numberString string = ""
	var stringToIntStringMap map[string]string = createStringToIntStringMap()
	fmt.Println(inputString)
	inputString = replaceStringWithStringMap(inputString, stringToIntStringMap)
	for _, char := range inputString {
		if unicode.IsDigit(char) {
			numberString += string(char)
		}
	}
	fmt.Println(inputString)
	fmt.Println(numberString)
	return numberString
}

func parseNumberToInt(inputString string) int {
	var intString string = string(inputString[0]) + string(inputString[len(inputString)-1])

	finalInt, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	return finalInt
}

func createStringToIntStringMap() map[string]string {
	var stringToIntStringMap map[string]string = make(map[string]string)

	stringToIntStringMap["zero"] = "z0o"
	stringToIntStringMap["one"] = "o1e"
	stringToIntStringMap["two"] = "t2o"
	stringToIntStringMap["three"] = "t3e"
	stringToIntStringMap["four"] = "f4r"
	stringToIntStringMap["five"] = "f5e"
	stringToIntStringMap["six"] = "s6x"
	stringToIntStringMap["seven"] = "s7n"
	stringToIntStringMap["eight"] = "e8t"
	stringToIntStringMap["nine"] = "n9n"

	return stringToIntStringMap
}

func replaceStringWithStringMap(inputString string, mapString map[string]string) string {
	for key, value := range mapString {
		inputString = strings.ReplaceAll(inputString, key, value)
	}
	return inputString
}
