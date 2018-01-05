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
	x1, v1, x2, v2 := parseInput(input)

	if (x1 < x2) && (v1 < v2) {
		fmt.Println("NO")
	} else {
		if (v1 != v2) && ((x2-x1)%(v1-v2)) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func parseInput(s string) (int, int, int, int) {
	nums := strings.Split(strings.Trim(s, "\n"), " ")

	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	c, _ := strconv.Atoi(nums[2])
	d, _ := strconv.Atoi(nums[3])

	return a, b, c, d
}
