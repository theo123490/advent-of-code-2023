package day2Tools

import (
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

type CubeBag struct {
	bagMap map[string]int
}
type Game struct {
	Id     int
	action []CubeAction
}

func ParseGame(inputString string) Game {
	var gameIdString string = inputString[:strings.IndexByte(inputString, ':')]
	var id int = getGameId(gameIdString)

	var actionArrayString string = inputString[strings.IndexByte(inputString, ':')+2:]
	var action []CubeAction = getActionArray(actionArrayString)

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

func CalculateActionCubeBag(game Game) CubeBag {
	var cubeMap map[string]int = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, action := range game.action {
		for _, cubeSet := range action.Bag {
			if cubeMap[cubeSet.Color] < cubeSet.amount {
				cubeMap[cubeSet.Color] = cubeSet.amount
			}
		}
	}

	return CubeBag{cubeMap}
}

func ExistingBagMap() CubeBag {
	var cubeBagMap map[string]int = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	return CubeBag{cubeBagMap}
}

func CheckIfBagMapPossible(existingBagMap CubeBag, currentBagMap CubeBag) bool {
	for color, _ := range existingBagMap.bagMap {
		if existingBagMap.bagMap[color] < currentBagMap.bagMap[color] {
			return false
		}
	}
	return true
}
