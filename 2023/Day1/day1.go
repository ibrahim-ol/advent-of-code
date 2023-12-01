package main

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Read input text file
	data, err := os.ReadFile("./input.txt")
	checkError(err)

	fmt.Println("Part 1: ", getPartOne(data))

	fmt.Println("Part 2: ", getPartTwo(data))

}
