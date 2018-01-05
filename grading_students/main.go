package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var input string

	_, scanErr := fmt.Scanln(&input)
	check(scanErr)

	count, convErr := strconv.Atoi(input)
	check(convErr)

	grades := make([]int, count)

	for i := 0; i < count; i++ {
		_, scanErr = fmt.Scanln(&input)
		check(scanErr)

		num, numErr := strconv.Atoi(input)
		check(numErr)

		grades[i] = num
	}

	for i := 0; i < len(grades); i++ {
		if grades[i] >= 38 && ((grades[i] + (5 - grades[i]%5) - grades[i]) < 3) {
			grades[i] += (5 - grades[i]%5)
		}
		fmt.Println(grades[i])
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln("Error:", e)
	}
}
