package day_one

import (
	"advent_2023/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func RunPartOne() int {
	input := utils.ReadFile("day_one/input.txt")
	inputRows := strings.Split(input, "\n")

	rxp := regexp.MustCompile("[1-9]")
	sum := 0

	for rowI := 0; rowI < len(inputRows); rowI++ {
		numbersInRow := rxp.FindAllString(inputRows[rowI], -1)

		if nil == numbersInRow {
			log.Fatal("unexpected input with no numbers in a row")
		}

		val, err := strconv.Atoi(numbersInRow[0] + numbersInRow[len(numbersInRow)-1])
		if nil != err {
			log.Fatal("unexpected error during string conversion to number", err)
		}

		sum += val
	}

	return sum
}

func RunPartTwo() int {
	input := utils.ReadFile("day_one/input.txt")
	fixOverlap(&input)

	inputRows := strings.Split(input, "\n")

	rxp := regexp.MustCompile("[1-9]|one|two|three|four|five|six|seven|eight|nine")
	sum := 0

	for rowI := 0; rowI < len(inputRows); rowI++ {
		numbersInRow := rxp.FindAllString(inputRows[rowI], -1)

		if nil == numbersInRow {
			log.Fatal("unexpected input with no numbers in a row")
		}

		firstVal := numbersInRow[0]
		secondVal := numbersInRow[len(numbersInRow)-1]
		val, err := strconv.Atoi(convertValue(firstVal) + convertValue(secondVal))
		if nil != err {
			log.Fatal("unexpected error during string conversion to number: ", err)
		}

		sum += val
	}

	return sum
}

// Wanted to try pointers
func fixOverlap(value *string) {
	// Stupid hardcoded edge case fix
	*value = strings.ReplaceAll(*value, "oneight", "oneeight")
	*value = strings.ReplaceAll(*value, "threeight", "threeeight")
	*value = strings.ReplaceAll(*value, "fiveight", "fiveeight")
	*value = strings.ReplaceAll(*value, "nineight", "nineeight")
	*value = strings.ReplaceAll(*value, "twone", "twoone")
	*value = strings.ReplaceAll(*value, "eightwo", "eighttwo")
}

var conversionMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func convertValue(value string) string {
	if 1 == len(value) {
		return value
	}
	return strings.ReplaceAll(value, value, conversionMap[value])
}
