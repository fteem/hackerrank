package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := parseNumber(reader.ReadString('\n'))
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = parseNumbers(reader.ReadString('\n'))
	}

	fmt.Println(sum(matrix))
}

func sum(matrix [][]int) float64 {
	diag1sum := 0
	diag2sum := 0

	for i := 0; i < len(matrix); i++ {
		diag1sum += matrix[i][i]
		diag2sum += matrix[i][len(matrix)-1-i]
	}

	return math.Abs(float64(diag1sum) - float64(diag2sum))
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
