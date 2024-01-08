package model

import (
	"log"
	"regexp"
	"strconv"
)

type Round struct {
	BallCount map[BallColor]int
}

func NewRound() *Round {
	return &Round{
		BallCount: map[BallColor]int{
			RED:   0,
			GREEN: 0,
			BLUE:  0,
		},
	}
}

func (r *Round) ParseRound(input string) {
	for _, color := range ColorsToCheck {
		match := regexp.MustCompile(colorParseMatchers[color]).FindStringSubmatch(input)
		if len(match) < 2 {
			continue
		}

		res, err := strconv.Atoi(match[1])
		if nil != err {
			log.Fatalf("could not convert color: '%v', in round: '%v', err: '%v' ", color, input, err)
		}

		r.BallCount[color] = res
	}
}
