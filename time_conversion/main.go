package main

import (
	"fmt"
	"log"
	"time"
)

const inputLayout = "03:04:05PM"
const outputLayout = "15:04:05"

func main() {
	var input string

	_, err := fmt.Scanln(&input)
	check(err)

	t, err := time.Parse(inputLayout, input)
	check(err)

	fmt.Println(t.Format(outputLayout))
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
