package main

import (
	"strings"
)

func part1(argument []byte) int {

	possibleSummation := 0
	expectedRedCount := 12
	expectedGreenCount := 13
	expectedBlueCount := 14

	for _, line := range strings.Split(string(argument), "\n") {
		firstSplit := strings.Split(line, ":")
		gameNumber := isNumber(strings.Split(firstSplit[0], " ")[1], 0)

		if gameNumber == 0 {
			panic("Game number should not be 0")
		}

		gameSets := strings.Split(firstSplit[1], ";")

		setIsValid := false
		for _, gameSet := range gameSets {
			red := extractColorNumber(gameSet, "red")
			green := extractColorNumber(gameSet, "green")
			blue := extractColorNumber(gameSet, "blue")

			if red <= expectedRedCount && green <= expectedGreenCount && blue <= expectedBlueCount {
				setIsValid = true
			} else {
				setIsValid = false
				break
			}
		}
		if setIsValid {
			possibleSummation += gameNumber
		}
	}

	return (possibleSummation)
}
