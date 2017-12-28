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

	n := parseNumber(reader.ReadString('\n'))
	numbers := parseNumbers(reader.ReadString('\n'))

	positive, negative, zeroes := fractions(numbers)
	fmt.Println(div(positive, n))
	fmt.Println(div(negative, n))
	fmt.Println(div(zeroes, n))
}

func div(a int, b int) float64 {
	return float64(a) / float64(b)
}

func fractions(numbers []int) (int, int, int) {
	positive := 0
	negative := 0
	zeroes := 0

	for i := 0; i < len(numbers); i++ {
		switch {
		case numbers[i] == 0:
			zeroes++
		case numbers[i] > 0:
			positive++
		case numbers[i] < 0:
			negative++
		}
	}
	return positive, negative, zeroes
}

func parseNumber(input string, err error) int {
	check(err)

	rawNumber := strings.Trim(input, "\n")
	number, _ := strconv.Atoi(rawNumber)

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

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}
