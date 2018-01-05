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
	n := parseNumber(input)

	input, err = reader.ReadString('\n')
	check(err)
	p := parseNumber(input)

	total := n / 2
	right := p / 2
	left := total - right

	if right < left {
		fmt.Println(right)
	} else {
		fmt.Println(left)
	}
}

func parseNumber(input string) int {
	number, _ := strconv.Atoi(strings.Trim(input, "\n"))
	return number
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}
