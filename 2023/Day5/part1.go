package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(argument []byte) int {

	seeds, dataMap := extractRanges(argument)

	// compute
	var seedLocations []int
	for _, seed := range seeds {
		valueToProcess := seed

		currentValue := passThroughTheProcess(valueToProcess, dataMap)

		seedLocations = append(seedLocations, currentValue)
	}
	return minInList(seedLocations)
}

func passThroughTheProcess(valueToProcess int, dataMap [][]RangeData) int {
	for _, process := range dataMap {
		valueToProcess = getDestinationFromSource(process, valueToProcess)
	}
	return valueToProcess
}

func getDestinationFromSource(ranges []RangeData, source int) int {

	for _, rangeData := range ranges {
		rangeStart := rangeData.sourceStart
		rangeEnd := rangeData.sourceStart + rangeData.rangeLength
		if source >= rangeStart && source <= rangeEnd {
			return rangeData.destinationStart + (source - rangeStart)
		}
	}

	return source
}

func minInList(list []int) int {
	min := list[0]
	for _, a := range list {
		if a < min {
			min = a
		}
	}
	return min
}
func extractRanges(lines []byte) ([]int, [][]RangeData) {
	seeds := []int{}
	dataMap := [][]RangeData{}
	var currentMapRange = -1

	argStrings := strings.Split(string(lines), "\n")

	for _, line := range argStrings {

		if strings.Trim(line, " ") == "" {
			continue
		}
		// get seeds
		if strings.HasPrefix(line, "seeds") {
			seeds = getNumbers(line)
			continue
		}

		// handle title line
		if strings.HasSuffix(strings.TrimSpace(line), ":") {
			currentMapRange += 1
			dataMap = append(dataMap, []RangeData{})
			continue
		}

		if currentMapRange != -1 {
			numbers := getNumbers(line)
			if len(numbers) == 3 {

				dataMap[currentMapRange] = append(dataMap[currentMapRange], createRangeData(numbers))
			}
		}

	}

	return seeds, dataMap
}
func createRangeData(numbers []int) RangeData {
	return RangeData{
		destinationStart: numbers[0],
		sourceStart:      numbers[1],
		rangeLength:      numbers[2],
	}
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

func findIndexOfTitle(hay []string, pin string) int {
	for i, hay := range hay {
		if hay == pin {
			return i
		}
	}
	return -1
}

type RangeData struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}
