package day_two

import (
	"advent_2023/day_two/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecutePartOne(t *testing.T) {
	expectedResult := 8

	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	limits := map[model.BallColor]int{
		model.RED:   12,
		model.GREEN: 13,
		model.BLUE:  14,
	}

	result := executePartOne(input, limits)

	assert.Equal(t, expectedResult, result, "result should match expected result")

}
func TestExecutePartTwo(t *testing.T) {
	expectedResult := 2286

	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	result := executePartTwo(input)

	assert.Equal(t, expectedResult, result, "result should match expected result")

}
