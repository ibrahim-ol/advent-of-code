package main

import (
	"strconv"
	"strings"
)

func extractColorNumber(argument string, color string) int {
	for _, cg := range strings.Split(argument, ",") {
		if strings.Contains(cg, color) {
			num := isNumber(strings.Replace(cg, color, "", 1), 0)
			if num != 0 {
				return num
			}
		}
	}
	return 0
}

func isNumber(argument string, deflt int) int {
	value, err := strconv.ParseInt(strings.Trim(argument, " "), 10, 64)
	if err != nil {
		return deflt
	} else {
		return int(value)
	}
}
