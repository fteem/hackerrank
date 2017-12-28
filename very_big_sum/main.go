package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	length := parseNumber(reader.ReadString('\n'))

	numbers := parseNumbers(reader.ReadString('\n'))

	sum := 0
	for i := 0; i < length; i++ {
		sum += numbers[i]
	}

	fmt.Print(sum)
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}

func parseNumber(input string, err error) int {
	check(err)

	rawNumber := strings.Trim(input, "\n")

	number, convErr := strconv.Atoi(rawNumber)
	check(convErr)

	return number
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
