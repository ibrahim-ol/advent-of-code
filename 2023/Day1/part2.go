package main

import (
	"strings"
)

func findFirstDigitMatch(argString string, reversed bool) int {
	if reversed {
		argString = reverseString(argString)
	}

	mapList := []NumMap{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}

	for index := range argString {

		isNum, num := isNumber(string(argString[index]))
		if isNum {
			return num
		}

		for _, mapL := range mapList {
			sliceEndIndex := index + len(mapL.key)
			if sliceEndIndex > len(argString) {
				continue
			}
			strSlice := argString[index:sliceEndIndex]
			compareKey := mapL.key
			if reversed {
				compareKey = reverseString(compareKey)
			}
			if strSlice == compareKey {
				return mapL.value
			}
		}
	}
	return 0
}

func getPartTwo(argument []byte) int {

	cleanCalibrationSum := 0
	for _, line := range strings.Split(string(argument), "\n") {
		cleanCalibrationSum += parseLine(line)
	}

	return cleanCalibrationSum

}

func parseLine(line string) int {
	firstDigit := findFirstDigitMatch(line, false)
	lastDigit := findFirstDigitMatch(line, true)

	fullValue := (firstDigit * 10) + lastDigit

	return int(fullValue)
}
