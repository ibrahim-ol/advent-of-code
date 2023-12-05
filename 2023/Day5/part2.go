package main

import (
	"fmt"
	"time"
)

func part2(argument []byte) int {
	startTime := time.Now()
	seeds, dataMap := extractRanges(argument)

	lowestSeedLocation := -1

	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedEnd := seedStart + seeds[i+1] - 1

		fmt.Println("Seed range ", seedStart, " to ", seedEnd)
		for j := seedStart; j <= seedEnd; j++ {
			valueToProcess := j
			currentValue := passThroughTheProcess(valueToProcess, dataMap)

			if lowestSeedLocation == -1 || currentValue < lowestSeedLocation {
				lowestSeedLocation = currentValue

				fmt.Println("================== Current Lowest location", lowestSeedLocation)
			}

		}
	}
	elapsed := time.Since(startTime)
	fmt.Println("Time taken ", elapsed)
	return lowestSeedLocation
}
