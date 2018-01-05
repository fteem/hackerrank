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
	row := strings.Split(strings.Trim(input, "\n"), " ")
	n, _ := strconv.Atoi(row[0])
	k, _ := strconv.Atoi(row[1])

	input, _ = reader.ReadString('\n')
	rawNumbers := strings.Split(strings.Trim(input, "\n"), " ")
	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		c, _ := strconv.Atoi(rawNumbers[i])
		numbers[i] = c
	}

	pairs := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i < j && (numbers[i]+numbers[j])%k == 0 {
				pairs++
			}
		}
	}

	fmt.Println(pairs)
}
