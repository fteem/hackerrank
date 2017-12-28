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
	numbers := parseInput(reader.ReadString('\n'))
	sort.Ints(numbers)

	fmt.Println(sum(numbers[:4]), sum(numbers[1:]))
}

func parseInput(input string, err error) []int {
	check(err)

	rawNumbers := strings.Split(strings.Trim(input, "\n"), " ")

	numbers := make([]int, len(rawNumbers))

	for i := 0; i < len(numbers); i++ {
		numbers[i], _ = strconv.Atoi(rawNumbers[i])
	}

	return numbers
}

func sum(s []int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}
