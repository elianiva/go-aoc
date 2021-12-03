package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"fmt"
	"os"
)

var days = map[string]func(){
	"1": day1.Solve,
	"2": day2.Solve,
	"3": day3.Solve,
}

func main() {
	day := os.Getenv("DAY")
	if day == "" {
		fmt.Println("DAY env variable wasn't provided.")
		os.Exit(1)
	}

	days[day]()
}
