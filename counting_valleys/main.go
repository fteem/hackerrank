package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	var err error

	_, err = fmt.Scanln(&input)
	check(err)
	stepsCount := parseNumber(input)

	_, err = fmt.Scanln(&input)
	check(err)
	steps := parseSteps(input)

	currentElevation := 0
	inValley := false
	valleys := 0

	for i := 0; i < stepsCount; i++ {
		switch steps[i] {
		case "U":
			currentElevation++
			if inValley && currentElevation >= 0 {
				inValley = false
				valleys++
			}
		case "D":
			currentElevation--
			if !inValley && currentElevation < 0 {
				inValley = true
			}
		}

	}

	fmt.Println(valleys)
}

func parseSteps(input string) []string {
	return strings.Split(input, "")
}

func parseNumber(input string) int {
	number, _ := strconv.Atoi(strings.Trim(input, "\n"))
	return number
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}
