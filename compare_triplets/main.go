package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type person struct {
	points  int
	numbers []int
}

func NewPerson(p int, n []int) person {
	return person{
		points:  p,
		numbers: n,
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	aliceNumbers := parseInput(reader.ReadString('\n'))
	alice := NewPerson(0, aliceNumbers)

	bobNumbers := parseInput(reader.ReadString('\n'))
	bob := NewPerson(0, bobNumbers)

	for i := 0; i < 3; i++ {
		if bob.numbers[i] > alice.numbers[i] {
			bob.points++
		} else if bob.numbers[i] < alice.numbers[i] {
			alice.points++
		}
	}

	fmt.Print(alice.points, bob.points)
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
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
