package d4

import (
	"advent_2023/utils"
	"regexp"
	"slices"
	"strings"
)

func RunPartOne() int {
	input := utils.ReadFile("days/d4/input.txt")
	input = strings.ReplaceAll(input, "\r\n", "\n")
	inputRows := strings.Split(input, "\n")

	return executePartOne(inputRows)
}

func executePartOne(inputRows []string) int {
	var numberList []int
	for _, row := range inputRows {
		numberGroups := regexp.MustCompile("[|:]").Split(row, -1)
		winningNumbers := regexp.MustCompile("\\d+\\d*").FindAllString(numberGroups[1], -1)
		numbersInHand := regexp.MustCompile("(\\d+\\d*)").FindAllString(numberGroups[2], -1)

		points := 0
		for _, number := range winningNumbers {
			if slices.Contains(numbersInHand, number) {
				if 0 == points {
					points = 1
					continue
				}
				points *= 2
			}
		}
		numberList = append(numberList, points)
	}

	sum := 0
	for _, val := range numberList {
		sum += val
	}

	return sum
}

//
//func RunPartTwo() int {
//	//input := utils.ReadFile("days/d3/input.txt")
//	//input = strings.ReplaceAll(input, "\r\n", "\n")
//	//inputRows := strings.Split(input, "\n")
//	//
//	//return executePartTwo(inputRows)
//}
