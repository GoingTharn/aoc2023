package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:2:1", y2023d2part1)
	registerSolution("2023:2:2", y2023d2part2)

}

type Game struct {
	red   int
	blue  int
	green int
}

func y2023d2part1(input string) string {
	myGame := Game{red: 12, green: 13, blue: 14}

	var validGames []int
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		gameNumber, pulls := parseLineInput(line)
		valid := isGameValid(myGame, pulls)

		if valid {
			validGames = append(validGames, gameNumber)
		}
	}

	sum := lo.Reduce(validGames, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)

	return fmt.Sprint(sum)
}

func parseLineInput(input string) (gameNumber int, pulls []string) {
	onColon := strings.Split(input, ":")
	gameData := onColon[0]
	gameStr := strings.TrimSpace(onColon[1])
	pulls = strings.Split(gameStr, ";")
	gameNumber, err := strconv.Atoi(strings.TrimPrefix(gameData, "Game "))
	if err != nil {
		panic(err)
	}
	return gameNumber, pulls
}

func compareGame(testGame Game, game Game) bool {
	return testGame.green >= game.green && testGame.red >= game.red && testGame.blue >= game.blue
}

func isGameValid(testGame Game, pulls []string) bool {
	for _, pull := range pulls {
		pullGame := parseGame(pull)
		valid := compareGame(testGame, pullGame)

		if !valid {
			return false
		}
	}
	return true
}

func setMaxBlocks(testGame Game, game Game) Game {
	if testGame.green < game.green {
		testGame.green = game.green
	}
	if testGame.red < game.red {
		testGame.red = game.red
	}
	if testGame.blue < game.blue {
		testGame.blue = game.blue
	}
	return testGame
}

func findMaxBlocks(testGame Game, pulls []string) Game {
	for _, pull := range pulls {
		pullGame := parseGame(pull)
		testGame = setMaxBlocks(testGame, pullGame)
	}
	return testGame
}

func parseGame(gameData string) Game {
	newGame := Game{}
	for _, elem := range strings.Split(gameData, ",") {
		splits := strings.Split(strings.TrimSpace(elem), " ")
		blockCount, _ := strconv.Atoi(splits[0])
		color := splits[1]
		switch color {
		case "red":
			newGame.red = blockCount
		case "blue":
			newGame.blue = blockCount
		case "green":
			newGame.green = blockCount
		}
	}
	return newGame
}

func y2023d2part2(input string) string {

	var validGames []int
	for _, line := range strings.Split(input, "\n") {
		myGame := Game{red: 0, green: 0, blue: 0}
		if len(line) == 0 {
			continue
		}
		_, pulls := parseLineInput(line)
		myGame = findMaxBlocks(myGame, pulls)
		validGames = append(validGames, myGame.blue*myGame.red*myGame.green)
	}

	sum := lo.Reduce(validGames, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)

	return fmt.Sprint(sum)
}
