package day_two

import (
	"advent_2023/day_two/part_one"
	"advent_2023/utils"
	"regexp"
	"strings"
)

func RunPartOne() int {
	input := utils.ReadFile("day_two/input.txt")
	inputRows := strings.Split(input, "\n")

	colorLimit := map[part_one.Color]int{
		part_one.RED:   12,
		part_one.GREEN: 13,
		part_one.BLUE:  14,
	}

	return executePartOne(inputRows, colorLimit)
}

func executePartOne(inputRows []string, limits map[part_one.Color]int) int {
	games := make([]part_one.Game, 0)

	for rowI := 0; rowI < len(inputRows); rowI++ {
		rowSlices := regexp.MustCompile("([:;])").Split(inputRows[rowI], -1)

		game := *part_one.NewGame()
		game.ParseGameNumber(rowSlices[0])

		// Skip first split, as it will include the game number
		// e.g. Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		for roundI := 1; roundI < len(rowSlices); roundI++ {
			round := *part_one.NewRound()
			round.ParseRound(rowSlices[roundI])

			game.Rounds = append(game.Rounds, round)
		}
		games = append(games, game)
	}

	sum := 0
	for _, game := range games {
		if game.IsGameWithinLimits(limits) {
			sum += game.RoundNumber
		}
	}

	return sum
}
