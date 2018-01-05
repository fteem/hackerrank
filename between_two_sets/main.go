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
	nA, nB := parseInput(input)

	setA := make([]int, nA)
	setB := make([]int, nB)

	input, _ = reader.ReadString('\n')
	rawA := strings.Split(strings.Trim(input, "\n"), " ")
	for i := 0; i < nA; i++ {
		num, _ := strconv.Atoi(rawA[i])
		setA[i] = num
	}

	input, _ = reader.ReadString('\n')
	rawB := strings.Split(strings.Trim(input, "\n"), " ")
	for i := 0; i < nB; i++ {
		num, _ := strconv.Atoi(rawB[i])
		setB[i] = num
	}

	for i := 0; i < len(setA); i++ {

	}
}

func parseInput(s string) (int, int) {
	nums := strings.Split(strings.Trim(s, "\n"), " ")

	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])

	return a, b
}
