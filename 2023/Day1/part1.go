package main

import (
	"strings"
)

func getFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		isNum, num := isNumber(string(line[i]))

		if isNum {
			return num
		}
	}
	return 0
}

func getLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		isNum, num := isNumber(string(line[i]))
		if isNum {
			return num
		}
	}
	return 0
}

func getPartOne(argument []byte) int {
	cleanCalibrationSum := 0
	for _, line := range strings.Split(string(argument), "\n") {
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)
		cleanCalibrationSum += int(firstDigit)*10 + int(lastDigit)
	}

	return (cleanCalibrationSum)
}
