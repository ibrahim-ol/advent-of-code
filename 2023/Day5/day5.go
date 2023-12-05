package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	sample, err := os.ReadFile("./sample.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Sample Part 1: ", part1([]byte(sample)))

	fmt.Println("Part 1: ", part1(data))

	fmt.Println("Sample Part 2: ", part2([]byte(sample)))

	fmt.Println("Part 2: ", part2(data))
}
