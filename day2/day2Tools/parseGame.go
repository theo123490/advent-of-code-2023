package day2Tools

import (
	"fmt"
	"strconv"
	"strings"
)

type Cube struct {
	Color string
}

type CubeBag struct {
	Bag map[Cube]int
}

type CubeAction struct {
	Bag map[Cube]int
}

type Game struct {
	id     int
	action []CubeAction
}

func ParseGame(inputString string) Game {
	var gameIdString string = inputString[:strings.IndexByte(inputString, ':')]
	var id int = getGameId(gameIdString)

	var actionArrayString string = inputString[strings.IndexByte(inputString, ':')+2:]

	fmt.Println(strconv.Itoa(id))
	fmt.Println(actionArrayString)
	return Game{}
}

func getGameId(gameIdString string) int {
	var idString string = gameIdString[strings.IndexByte(gameIdString, ' ')+1:]
	var id, err = strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	return id
}

func parseIdString(IdString string) int {
	return 0
}

func parseGameAction(ActionString string) CubeAction {
	return CubeAction{}
}
