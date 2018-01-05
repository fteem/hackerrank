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
	length, _ := strconv.Atoi(strings.Trim(input, "\n"))

	input, _ = reader.ReadString('\n')
	scores := strings.Split(strings.Trim(input, "\n"), " ")

	brokenBest := 0
	brokenWorst := 0

	best := 0
	worst := 0
	for i := 0; i < length; i++ {
		currentScore, _ := strconv.Atoi(scores[i])
		bestScore, _ := strconv.Atoi(scores[best])
		worstScore, _ := strconv.Atoi(scores[worst])

		if currentScore > bestScore {
			best = i
			brokenBest++
		}
		if currentScore < worstScore {
			worst = i
			brokenWorst++
		}
	}

	fmt.Println(brokenBest, brokenWorst)
}
