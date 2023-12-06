package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// buttonSecHeld * (TotalAllowedTime - buttonSecHeld) = DistanceTravelled
func distanceTravelled(secHeld, allowed int) int {
	return secHeld * (allowed - secHeld)
}

// (B * T) - (B * B) = D
// BT - BB = D
// BB - BT + D = 0
// b±√(b²-4ac))/(2a)
// a = 1
// b = -T
// c = D
func timeUsedForDistance(distance, totalTime int) int {
	determinant := (math.Pow(float64(totalTime), 2.0)) - float64(4*distance)

	sol1 := 0.5 * (float64(totalTime) + math.Sqrt(determinant))
	sol2 := 0.5 * (float64(totalTime) - math.Sqrt(determinant))

	return int(math.Floor(math.Min(sol1, sol2)))
}

// 0 1 2 3 4 5 6 7
// 3 4
// 2 5
// 1 6
// 0 7

// 0 1 2 3 4 5 6 7 8
// 4
// 3 5
// 2 6
// 1 7
// 0 8

// betterNoOfTimes = total time - ((cHP + 1) * 2) + 1

func part1(argument []byte) int {

	argStrings := strings.Split(string(argument), "\n")

	times := getNumbers(argStrings[0])
	distances := getNumbers(argStrings[1])

	// starting from 1 (because multiplication)
	result := 1
	for index, time := range times {

		timeHeldToGetDistance := timeUsedForDistance(distances[index], time)

		betterNoOfTimes := time - ((timeHeldToGetDistance + 1) * 2) + 1

		result *= betterNoOfTimes

	}

	return result
}

func getNumbers(line string) []int {
	re := regexp.MustCompile(`(\d+)`)
	var nums []int
	matches := re.FindAll([]byte(line), -1)

	for _, w := range matches {
		num, err := toNumber(string(w))
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func toNumber(argument string) (int, error) {
	value, err := strconv.ParseInt(strings.Trim(argument, " "), 10, 64)
	if err != nil {
		return 0, err
	} else {
		return int(value), nil
	}
}
