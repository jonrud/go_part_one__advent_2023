package model

type BallColor string

const (
	RED   = "RED"
	GREEN = "GREEN"
	BLUE  = "BLUE"
)

var ColorsToCheck = []BallColor{RED, GREEN, BLUE}

// Not using constants in regexp on purpose
var colorParseMatchers = map[BallColor]string{
	RED:   "([0-9]*) red",
	GREEN: "([0-9]*) green",
	BLUE:  "([0-9]*) blue",
}
