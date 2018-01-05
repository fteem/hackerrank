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

	reader := bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')
	check(err)
	total := parseNumber(input)

	input, err = reader.ReadString('\n')
	check(err)
	colors := parseNumbers(input)

	socksPerColor := make(map[int]int)
	for i := 0; i < total; i++ {
		socksPerColor[colors[i]]++
	}

	pairs := 0
	for _, count := range socksPerColor {
		pairs += count / 2
	}

	fmt.Println(pairs)
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
