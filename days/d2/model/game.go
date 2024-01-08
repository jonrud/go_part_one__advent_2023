package model

import (
	"log"
	"regexp"
	"strconv"
)

type Game struct {
	Rounds      []Round
	RoundNumber int
}

func NewGame() *Game {
	return &Game{
		Rounds:      []Round{},
		RoundNumber: 0,
	}
}

func (g *Game) CalculateMaxCounts() map[BallColor]int {
	maxCounts := map[BallColor]int{}

	for _, color := range ColorsToCheck {
		maxCounts[color] = g.GetMaxColorCount(color)
	}

	return maxCounts
}

func (g *Game) IsGameWithinLimits(limits map[BallColor]int) bool {
	for _, color := range ColorsToCheck {
		if limits[color] < g.GetMaxColorCount(color) {
			return false
		}
	}

	return true
}

func (g *Game) GetMaxColorCount(color BallColor) int {
	// Not very efficient, but cba bulking it

	maxColorCount := 0
	for _, round := range g.Rounds {
		if round.BallCount[color] > maxColorCount {
			maxColorCount = round.BallCount[color]
		}
	}

	return maxColorCount
}

func (g *Game) ParseGameNumber(input string) {
	match := regexp.MustCompile("Game ([0-9]*)").FindStringSubmatch(input)
	if len(match) < 2 {
		log.Fatalf("could not find game number in input: '%v'", input)
	}

	res, err := strconv.Atoi(match[1])
	if nil != err {
		log.Fatalf("could not convert round number, value: '%v', err: '%v'", match[1], err)
	}

	g.RoundNumber = res
}
