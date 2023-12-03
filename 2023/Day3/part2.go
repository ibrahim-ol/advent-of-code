package main

import (
	"regexp"
	"strings"
)

func getStarPosOnLine(line string) []int {
	re, _ := regexp.Compile(`(\*)`)
	var response []int
	matches := re.FindAllIndex([]byte(line), -1)
	for _, match := range matches {
		response = append(response, match[0])
	}
	return response
}

func getNumberPosOnMultiLine(lines []string) []NumPos {
	var res []NumPos
	for _, line := range lines {
		resp := getNumberPosOnLine(line)
		res = append(res, resp...)
	}
	return res
}

func isIntersecting(start int, end int, star int) bool {
	startIndex := star - 1
	endIndex := star + 1

	intersectionLogic := (start >= startIndex && start <= endIndex) || (end >= startIndex && end <= endIndex)
	if intersectionLogic {
		return true
	}

	return false
}

func part2(argument []byte) int {

	argStrings := strings.Split(string(argument), "\n")
	argStringsLength := len(argStrings)
	var gearRatioSum int
	for lineIndex, line := range argStrings {
		starPosition := getStarPosOnLine(line)

		var surroundingLines []string
		// // add previous line if we are not on the first line
		if lineIndex != 0 {
			surroundingLines = append(surroundingLines, argStrings[lineIndex-1])
		}
		// // add the line itself
		surroundingLines = append(surroundingLines, line)
		// add the next line if we are not on the last line
		if lineIndex != argStringsLength-1 {
			surroundingLines = append(surroundingLines, argStrings[lineIndex+1])
		}

		for _, position := range starPosition {
			numPos := getNumberPosOnMultiLine(surroundingLines)

			var adjNum []int

			for _, num := range numPos {
				intersect := isIntersecting(num.startIndex, num.endIndex-1, position)
				if intersect {
					adjNum = append(adjNum, toNumber(num.value, 0))
				}

			}
			if len(adjNum) == 2 {
				gearRatioSum += adjNum[0] * adjNum[1]
			}
		}

	}

	return gearRatioSum
}
