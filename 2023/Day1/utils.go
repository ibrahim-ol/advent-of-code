package main

import "strconv"

type NumMap = struct {
	key   string
	value int
}

func reverseString(argString string) string {
	reversed := ""
	for index := range argString {
		reversed += string(argString[len(argString)-index-1])
	}
	return reversed
}

func isNumber(argument string) (bool, int) {
	value, err := strconv.ParseInt(argument, 10, 64)
	if err != nil {
		return false, 0
	} else {
		return true, int(value)
	}
}
