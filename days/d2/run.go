package d2

import (
	model2 "advent_2023/days/d2/model"
	"advent_2023/utils"
	"regexp"
	"strings"
)

func RunPartOne() int {
	input := utils.ReadFile("days/d2/input.txt")
	inputRows := strings.Split(input, "\n")

	colorLimit := map[model2.BallColor]int{
		model2.RED:   12,
		model2.GREEN: 13,
		model2.BLUE:  14,
	}

	return executePartOne(inputRows, colorLimit)
}

func executePartOne(inputRows []string, limits map[model2.BallColor]int) int {
	games := make([]model2.Game, 0)

	for rowI := 0; rowI < len(inputRows); rowI++ {
		rowSlices := regexp.MustCompile("([:;])").Split(inputRows[rowI], -1)

		game := *model2.NewGame()
		game.ParseGameNumber(rowSlices[0])

		// Skip first split, as it will include the game number
		// e.g. Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		for roundI := 1; roundI < len(rowSlices); roundI++ {
			round := *model2.NewRound()
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
	input := utils.ReadFile("days/d2/input.txt")
	inputRows := strings.Split(input, "\n")

	return executePartTwo(inputRows)
}

func executePartTwo(inputRows []string) int {
	games := make([]model2.Game, 0)

	for rowI := 0; rowI < len(inputRows); rowI++ {
		rowSlices := regexp.MustCompile("([:;])").Split(inputRows[rowI], -1)

		game := *model2.NewGame()
		game.ParseGameNumber(rowSlices[0])

		// Skip first slice, as it will include the game number
		// e.g. Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		for roundI := 1; roundI < len(rowSlices); roundI++ {
			round := *model2.NewRound()
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
