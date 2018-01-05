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
	var input string
	var err error

	reader := bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')
	check(err)

	row := parseNumbers(input)
	budget := row[0]
	keyboardsCount := row[1]
	drivesCount := row[2]

	input, err = reader.ReadString('\n')
	check(err)
	keyboards := parseNumbers(input)
	sort.Sort(sort.Reverse(sort.IntSlice(keyboards)))

	input, err = reader.ReadString('\n')
	check(err)
	drives := parseNumbers(input)
	sort.Sort(sort.IntSlice(drives))

	max := -1
	for i, j := 0, 0; i < keyboardsCount; i++ {
		for ; j < drivesCount; j++ {
			sum := drives[j] + keyboards[i]
			if sum > budget {
				break
			}
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
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
