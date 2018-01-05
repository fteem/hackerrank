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

	// Get s & t
	input, _ = reader.ReadString('\n')
	s, t := getPair(input)

	// Get a & b
	input, _ = reader.ReadString('\n')
	a, b := getPair(input)

	// Get m & n
	input, _ = reader.ReadString('\n')
	m, n := getPair(input)

	// Get fallen oranges
	oDistances := make([]int, m)
	input, _ = reader.ReadString('\n')
	oDistancesRaw := strings.Split(strings.Trim(input, "\n"), " ")
	for i := 0; i < m; i++ {
		oDistances[i], _ = strconv.Atoi(oDistancesRaw[i])
	}

	// Get fallen apples
	aDistances := make([]int, n)
	input, _ = reader.ReadString('\n')
	aDistancesRaw := strings.Split(strings.Trim(input, "\n"), " ")
	for i := 0; i < n; i++ {
		aDistances[i], _ = strconv.Atoi(aDistancesRaw[i])
	}

	oranges := 0
	for i := 0; i < m; i++ {
		distance := a + oDistances[i]
		if distance >= s && distance <= t {
			oranges++
		}
	}
	fmt.Println(oranges)

	apples := 0
	for i := 0; i < n; i++ {
		distance := b + aDistances[i]
		if distance >= s && distance <= t {
			apples++
		}
	}
	fmt.Println(apples)
}

func getPair(i string) (int, int) {
	parts := strings.Split(strings.Trim(i, "\n"), " ")
	p1, _ := strconv.Atoi(parts[0])
	p2, _ := strconv.Atoi(parts[1])

	return p1, p2
}
