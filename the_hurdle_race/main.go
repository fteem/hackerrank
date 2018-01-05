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
	nums := parseNumbers(input)

	jump := nums[1]

	input, err = reader.ReadString('\n')
	check(err)
	hurdles := parseNumbers(input)
	sort.Ints(hurdles)

	tallestHurdle := hurdles[len(hurdles)-1]
	if jump >= tallestHurdle {
		fmt.Println(0)
	} else {
		fmt.Println(tallestHurdle - jump)
	}
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
