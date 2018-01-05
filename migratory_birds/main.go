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

	reader := bufio.NewReader(os.Stdin)

	input, _ = reader.ReadString('\n')
	n := parseNumber(input)

	input, _ = reader.ReadString('\n')
	types := parseNumbers(input, n)

	occurances := make(map[int]int)

	for i := 0; i < n; i++ {
		occurances[types[i]]++
	}

	max := 0
	for i := 0; i < n; i++ {
		if occurances[max] < occurances[i] {
			max = i
		}
	}
	fmt.Println(max)
}

func parseNumber(input string) int {
	number, _ := strconv.Atoi(strings.Trim(input, "\n"))
	return number
}

func parseNumbers(input string, length int) []int {
	rawNumbers := strings.Split(strings.Trim(input, "\n"), " ")
	numbers := make([]int, length)

	for i := 0; i < length; i++ {
		c, _ := strconv.Atoi(rawNumbers[i])
		numbers[i] = c
	}

	return numbers
}
