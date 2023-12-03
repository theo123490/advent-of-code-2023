package day2Tools

import (
	"fmt"
	"strconv"
	"strings"
)

type CubeSet struct {
	Color  string
	amount int
}

type CubeAction struct {
	Bag []CubeSet
}

type Game struct {
	id     int
	action []CubeAction
}

func ParseGame(inputString string) Game {
	var gameIdString string = inputString[:strings.IndexByte(inputString, ':')]
	var id int = getGameId(gameIdString)

	var actionArrayString string = inputString[strings.IndexByte(inputString, ':')+2:]
	var action []CubeAction = getActionArray(actionArrayString)

	fmt.Println(strconv.Itoa(id))
	fmt.Println(action)

	var game Game = Game{id, action}
	return game
}

func getGameId(gameIdString string) int {
	var idString string = gameIdString[strings.IndexByte(gameIdString, ' ')+1:]
	var id, err = strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	return id
}

func getActionArray(actionArrayString string) []CubeAction {
	var fullActionStringArray []string = strings.Split(actionArrayString, ";")
	var actionArray []CubeAction = make([]CubeAction, 0)
	for _, fullActionString := range fullActionStringArray {
		actionArray = append(actionArray, createCubeAction(fullActionString))
	}
	return actionArray
}

func createCubeAction(fullActionString string) CubeAction {
	var actionStringArray []string = strings.Split(fullActionString, ",")
	var cubeAction CubeAction = CubeAction{}
	for _, actionString := range actionStringArray {
		var cubeSet CubeSet = stringToCube(actionString)
		cubeAction.Bag = append(cubeAction.Bag, cubeSet)
	}
	return cubeAction
}

func stringToCube(CubeString string) CubeSet {
	var cubeStringField []string = strings.Fields(CubeString)
	var cubeColor string = cubeStringField[1]
	cubeAmount, err := strconv.Atoi(cubeStringField[0])
	if err != nil {
		panic(err)
	}

	var cubeSet CubeSet = CubeSet{cubeColor, cubeAmount}
	return cubeSet
}
