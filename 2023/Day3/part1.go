package main

import (
	"regexp"
	"strconv"
	"strings"
)

type NumPos struct {
	startIndex int
	endIndex   int
	value      string
}

func getNumberPosOnLine(line string) []NumPos {
	 re, _ := regexp.Compile(`(\d+)`)
	var res []NumPos
	resp := re.FindAllIndex([]byte(line), -1)
	for _, num := range resp {
		res = append(res, NumPos{startIndex: num[0], endIndex: num[1], value: line[num[0]:num[1]]})
	}
	return res
}


func hasAdjacentSymbol(around string) bool {
	hasAdjacent := false
	for _, chr := range around {
		if strings.Contains("1234567890.", string(chr)) {
			continue
		}
		hasAdjacent = true
		break
	}
	return hasAdjacent
}

func getSurrounding(around []string, startIndex int, endIndex int) string {
	aStartIndex := startIndex - 1
	aEndIndex := endIndex + 1
	if startIndex == 0 {
		aStartIndex = startIndex
	}
	if endIndex >= len(around[0]) {
		aEndIndex = endIndex
	}

	var res []string
	for _, line := range around {
		line = strings.Trim(line, " ")

		res = append(res, string(line[aStartIndex:aEndIndex]))
	}
	return strings.Join(res, "")
}

func part1(argument []byte) int {

	var numWithAdjacentSum int
	argStrings := strings.Split(string(argument), "\n")
	argStringsLength := len(argStrings)
	for lineIndex, line := range argStrings {
		numPositions := getNumberPosOnLine(line)

		var surrounding []string
		// add previous line if we are not on the first line
		if lineIndex != 0 {
			surrounding = append(surrounding, argStrings[lineIndex-1])
		}
		// add the line itself
		surrounding = append(surrounding, line)
		// add the next line if we are not on the last line
		if lineIndex != argStringsLength-1 {
			surrounding = append(surrounding, argStrings[lineIndex+1])
		}

		for _, position := range numPositions {
			if hasAdjacentSymbol(getSurrounding((surrounding), position.startIndex, position.endIndex)) {
				numWithAdjacentSum += toNumber(position.value, 0)
			}
		}
	}

	return numWithAdjacentSum
}

func toNumber(argument string, deflt int) int {
	value, err := strconv.ParseInt(strings.Trim(argument, " "), 10, 64)
	if err != nil {
		return deflt
	} else {
		return int(value)
	}
}
