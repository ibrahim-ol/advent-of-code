package main

import (
	"strings"
)

func part2(argument []byte) int {

	argStrings := strings.Split(string(argument), "\n")

	extraCardMap := map[int]int{}
	for index, line := range argStrings {
		cardNumber := index + 1
		winning, bucket := process_line(line)
		occurence := 0

		for _, num := range bucket {
			for _, win := range winning {
				if num == win {
					occurence++
				}
			}
		}

		for i := 1; i <= occurence; i++ {
			extraCardMap[cardNumber+i] += 1 + extraCardMap[cardNumber]
		}
	}

	extras := 0
	for _, v := range extraCardMap {
		extras += v
	}
	return len(argStrings) + extras
}
