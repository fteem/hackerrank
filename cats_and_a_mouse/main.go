package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	var err error

	_, err = fmt.Scanln(&input)
	check(err)
	queries := parseNumber(input)

	scanner := bufio.NewScanner(os.Stdin)

	row := 1
	for scanner.Scan() {
		positions := parseNumbers(scanner.Text())
		catA := positions[0]
		catB := positions[1]
		mouseC := positions[2]

		if abs(mouseC-catA) < abs(mouseC-catB) {
			fmt.Println("Cat A")
		} else if abs(mouseC-catA) > abs(mouseC-catB) {
			fmt.Println("Cat B")
		} else {
			fmt.Println("Mouse C")
		}

		if row == queries {
			break
		}
		row++
	}

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parseNumber(input string) int {
	number, _ := strconv.Atoi(strings.Trim(input, "\n"))
	return number
}

func parseNumbers(input string) []int {
	rawNumbers := strings.Split(strings.Trim(input, "\n"), " ")
	numbers := make([]int, len(rawNumbers))

	for i := 0; i < len(numbers); i++ {
		c, _ := strconv.Atoi(rawNumbers[i])
		numbers[i] = c
	}

	return numbers
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}
