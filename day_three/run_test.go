package day_three

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecutePartOne(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expectedResult := 4361

	result := executePartOne(input)

	assert.Equal(t, expectedResult, result, "result should match expected result")
}
func TestExecutePartTwo(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expectedResult := 467835

	result := executePartTwo(input)

	assert.Equal(t, expectedResult, result, "result should match expected result")
}
