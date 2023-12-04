package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func process_line(line string) ([]int, []int) {
	re := regexp.MustCompile(`(\d+)`)

	nums := strings.Split(line, ":")[1]
	wSide := strings.Split(nums, " | ")[0]
	mSide := strings.Split(nums, " | ")[1]
	var wNums []int
	var bNums []int
	wResp := re.FindAll([]byte(wSide), -1)
	bResp := re.FindAll([]byte(mSide), -1)

	for _, w := range wResp {
		wNums = append(wNums, toNumber(string(w), 0))
	}
	for _, b := range bResp {
		bNums = append(bNums, toNumber(string(b), 0))
	}
	return wNums, bNums

}

func part1(argument []byte) int {

	argStrings := strings.Split(string(argument), "\n")
	totalPoint := 0
	for _, line := range argStrings {
		winning, bucket := process_line(line)
		occurence := 0
		for _, num := range bucket {
			for _, win := range winning {
				if num == win {
					occurence++
				}
			}
		}

		totalPoint += int(math.Pow(2, float64(occurence)-1))
	}
	return totalPoint
}

func toNumber(argument string, deflt int) int {
	value, err := strconv.ParseInt(strings.Trim(argument, " "), 10, 64)
	if err != nil {
		return deflt
	} else {
		return int(value)
	}
}
