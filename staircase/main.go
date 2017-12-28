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

	for i := 1; i <= n; i++ {
		spaces := n - i
		pounds := i
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < pounds; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func parseNumber(input string, err error) int {
	check(err)

	rawNumber := strings.Trim(input, "\n")
	number, _ := strconv.Atoi(rawNumber)

	return number
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}
