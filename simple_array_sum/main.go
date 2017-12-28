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

	rawLength, e1 := reader.ReadString('\n')
	check(e1)

	rawCollection, e2 := reader.ReadString('\n')
	check(e2)

	_ = stringToInt(strings.Trim(rawLength, "\n"))

	collection := stringsToInts(strings.Split(strings.Trim(rawCollection, "\n"), " "))

	fmt.Println(sum(collection))
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}

func sum(ints []int) int {
	sum := 0
	for i := 0; i < len(ints); i++ {
		sum += ints[i]
	}

	return sum
}

func stringToInt(s string) int {
	num, err := strconv.Atoi(s)
	check(err)

	return num
}

func stringsToInts(s []string) []int {
	ints := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		ints[i] = stringToInt(strings.Trim(s[i], "\n"))
	}

	return ints
}
