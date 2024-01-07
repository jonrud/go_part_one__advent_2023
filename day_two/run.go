package day_two

import (
	"advent_2023/day_two/model"
	"advent_2023/utils"
	"regexp"
	"strings"
)

func RunPartOne() int {
	input := utils.ReadFile("day_two/input.txt")
	inputRows := strings.Split(input, "\n")

	colorLimit := map[model.BallColor]int{
		model.RED:   12,
		model.GREEN: 13,
		model.BLUE:  14,
	}

	return executePartOne(inputRows, colorLimit)
}

func executePartOne(inputRows []string, limits map[model.BallColor]int) int {
	games := make([]model.Game, 0)

	for rowI := 0; rowI < len(inputRows); rowI++ {
		rowSlices := regexp.MustCompile("([:;])").Split(inputRows[rowI], -1)

		game := *model.NewGame()
		game.ParseGameNumber(rowSlices[0])

		// Skip first split, as it will include the game number
		// e.g. Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		for roundI := 1; roundI < len(rowSlices); roundI++ {
			round := *model.NewRound()
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

func RunPartTwo() int {
	input := utils.ReadFile("day_two/input.txt")
	inputRows := strings.Split(input, "\n")

	return executePartTwo(inputRows)
}

func executePartTwo(inputRows []string) int {
	games := make([]model.Game, 0)

	for rowI := 0; rowI < len(inputRows); rowI++ {
		rowSlices := regexp.MustCompile("([:;])").Split(inputRows[rowI], -1)

		game := *model.NewGame()
		game.ParseGameNumber(rowSlices[0])

		// Skip first slice, as it will include the game number
		// e.g. Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		for roundI := 1; roundI < len(rowSlices); roundI++ {
			round := *model.NewRound()
			round.ParseRound(rowSlices[roundI])

			game.Rounds = append(game.Rounds, round)
		}

		games = append(games, game)
	}

	sum := 0
	for _, game := range games {
		maxCounts := game.CalculateMaxCounts()

		localSum := -1
		for _, count := range maxCounts {
			if -1 == localSum {
				localSum = count
				continue
			}
			localSum *= count
		}
		sum += localSum
	}

	return sum
}
