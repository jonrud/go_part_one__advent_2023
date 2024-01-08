package model

import (
	"log"
	"strconv"
)

type Part struct {
	Value        string
	IsEnginePart bool
}

func NewPart() *Part {
	return &Part{
		Value:        "",
		IsEnginePart: false,
	}
}

func (p Part) GetPartValue() int {
	val, err := strconv.Atoi(p.Value)
	if err != nil {
		log.Fatalf("part value '%v' is not a number", p.Value)
	}
	return val
}
