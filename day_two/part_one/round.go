package part_one

import (
	"log"
	"regexp"
	"strconv"
)

type Color string

const (
	RED   = "RED"
	GREEN = "GREEN"
	BLUE  = "BLUE"
)

var ColorsToCheck = []Color{RED, GREEN, BLUE}

// Not using constants in regexp on purpose
var colorParseMatchers = map[Color]string{
	RED:   "([0-9]*) red",
	GREEN: "([0-9]*) green",
	BLUE:  "([0-9]*) blue",
}

type Round struct {
	BallCount map[Color]int
}

func NewRound() *Round {
	return &Round{
		BallCount: map[Color]int{},
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
