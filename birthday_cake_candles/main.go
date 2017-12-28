package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	candles := parseNumber(reader.ReadString('\n'))
	heights := parseNumbers(reader.ReadString('\n'))
	sort.Ints(heights)

	_ = candles
	maxHeight := heights[len(heights)-1]

	fmt.Println(occurances(heights, maxHeight))
}

func occurances(s []int, n int) int {
	occ := 0

	for i := 0; i < len(s); i++ {
		if s[i] == n {
			occ++
		}
	}

	return occ
}

func parseNumbers(input string, err error) []int {
	check(err)

	rawNumbers := strings.Split(strings.Trim(input, "\n"), " ")

	numbers := make([]int, len(rawNumbers))

	for i := 0; i < len(numbers); i++ {
		numbers[i], _ = strconv.Atoi(rawNumbers[i])
	}

	return numbers
}

func parseNumber(input string, err error) int {
	check(err)

	rawNumber := strings.Trim(input, "\n")
	number, _ := strconv.Atoi(rawNumber)

	return number
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}
