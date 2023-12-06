package main

import (
	"strconv"
	"strings"
)

func concatenateNums(nums []int) int {
	var concat = []string{}
	for _, num := range nums {
		concat = append(concat, strconv.Itoa(num))
	}
	num, _ := toNumber(strings.Join(concat, ""))
	return num
}

func part2(argument []byte) int {

	argStrings := strings.Split(string(argument), "\n")

	times := []int{concatenateNums(getNumbers(argStrings[0]))}
	distances := []int{concatenateNums(getNumbers(argStrings[1]))}

	result := 1
	for index, time := range times {

		timeHeldToGetDistance := timeUsedForDistance(distances[index], time)

		betterNoOfTimes := time - ((timeHeldToGetDistance + 1) * 2) + 1

		result *= betterNoOfTimes
	}

	return result
}
