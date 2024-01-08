package day_three

import (
	model2 "advent_2023/day_three/model"
	"advent_2023/utils"
	"strings"
)

func RunPartOne() int {
	input := utils.ReadFile("day_three/input.txt")
	input = strings.ReplaceAll(input, "\r\n", "\n")
	inputRows := strings.Split(input, "\n")

	return executePartOne(inputRows)
}

func executePartOne(inputRows []string) int {
	rowCount := len(inputRows)
	columnCount := len(inputRows[0])

	referenceMap := model2.NewReferenceMap(rowCount, columnCount)
	parts := populateParts(inputRows, rowCount, columnCount, &referenceMap)
	populateSymbols(inputRows, rowCount, columnCount, &referenceMap)

	sum := sumEngineParts(parts)

	return sum
}

func RunPartTwo() int {
	input := utils.ReadFile("day_three/input.txt")
	input = strings.ReplaceAll(input, "\r\n", "\n")
	inputRows := strings.Split(input, "\n")

	return executePartTwo(inputRows)
}

func executePartTwo(inputRows []string) int {
	rowCount := len(inputRows)
	columnCount := len(inputRows[0])

	referenceMap := model2.NewReferenceMap(rowCount, columnCount)
	populateParts(inputRows, rowCount, columnCount, &referenceMap)
	populateGears(inputRows, rowCount, columnCount, &referenceMap)

	sum := sumGearMultiplier(referenceMap)

	return sum
}

func populateParts(inputRows []string, rowCount int, columnCount int, referenceMap *model2.ReferenceMap) []*model2.Part {
	var parts []*model2.Part
	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		row := strings.Split(inputRows[rowIndex], "")

		var part *model2.Part = nil
		for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
			if !utils.IsNumber(row[columnIndex]) {
				// If not a number = always skip
				continue
			}

			if part == nil {
				part = model2.NewPart()
			}
			part.Value += row[columnIndex]
			referenceMap.Parts[rowIndex][columnIndex] = part

			isEndOfColumn := columnIndex == referenceMap.ColumnMax
			isEndOfSequence := !isEndOfColumn && !utils.IsNumber(row[columnIndex+1])
			if isEndOfColumn || isEndOfSequence {
				parts = append(parts, part)
				part = nil
				// finish part
			}
			continue

		}
	}
	return parts
}

func populateSymbols(inputRows []string, rowCount int, columnCount int, referenceMap *model2.ReferenceMap) {
	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		row := strings.Split(inputRows[rowIndex], "")
		for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
			isCurrentNumber := utils.IsNumber(row[columnIndex])
			isCurrentDot := row[columnIndex] == "."

			if !isCurrentNumber && !isCurrentDot {
				// If it's not a number or a dot - it's a symbol
				referenceMap.CheckForPartsNearSymbol(rowIndex, columnIndex)
			}
		}
	}
}

func sumEngineParts(parts []*model2.Part) int {
	sum := 0
	for _, part := range parts {
		if part.IsEnginePart {
			sum += part.GetPartValue()
		}
	}
	return sum
}

func populateGears(inputRows []string, rowCount int, columnCount int, referenceMap *model2.ReferenceMap) {
	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		row := strings.Split(inputRows[rowIndex], "")
		for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
			isGear := row[columnIndex] == "*"

			if isGear {
				// If it's not a number or a dot - it's a symbol
				referenceMap.CheckForPartsNearSymbol(rowIndex, columnIndex)
			}
		}
	}
}

func sumGearMultiplier(referenceMap model2.ReferenceMap) int {
	sum := 0
	for _, gearMultiplier := range referenceMap.GearMultipliers {
		sum += gearMultiplier
	}
	return sum
}
