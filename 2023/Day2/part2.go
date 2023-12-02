package main

import (
	"strings"
)

func part2(argument []byte) int {

	summation := 0

	for _, line := range strings.Split(string(argument), "\n") {
		firstSplit := strings.Split(line, ":")
		gameNumber := isNumber(strings.Split(firstSplit[0], " ")[1], 0)

		if gameNumber == 0 {
			panic("Game number should not be 0")
		}

		gameSets := strings.Split(firstSplit[1], ";")

		currentRedMax := 0
		currentGreenMax := 0
		currentBlueMax := 0
		for _, gameSet := range gameSets {
			red := extractColorNumber(gameSet, "red")
			green := extractColorNumber(gameSet, "green")
			blue := extractColorNumber(gameSet, "blue")

			if red > currentRedMax {
				currentRedMax = red
			}

			if green > currentGreenMax {
				currentGreenMax = green
			}

			if blue > currentBlueMax {
				currentBlueMax = blue
			}

		}
		power := currentBlueMax *
			currentGreenMax * currentRedMax

		summation += power

	}

	return (summation)
}
