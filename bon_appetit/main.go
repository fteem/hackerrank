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
	nums := parseNumbers(input)
	itemsOrdered := nums[0]
	skippedItem := nums[1]

	input, err = reader.ReadString('\n')
	check(err)
	costs := parseNumbers(input)

	input, err = reader.ReadString('\n')
	check(err)
	charged := parseNumber(input)

	costOfSharedMeals := 0
	for i := 0; i < itemsOrdered; i++ {
		if i != skippedItem {
			costOfSharedMeals += costs[i]
		}
	}

	if costOfSharedMeals/2 == charged {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(charged - costOfSharedMeals/2)
	}
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
