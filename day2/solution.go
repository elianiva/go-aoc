package day2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("======== DAY 2 ========")

	inputFile, err := ioutil.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	input := linesToTuple(inputFile)

	pos := findPos(input)
	posWithAim := findPosWithAim(input)

	fmt.Println("PART 1:", pos)
	fmt.Println("PART 2:", posWithAim)
}

func linesToTuple(str []byte) [][]string {
	lines := strings.Split(string(str), "\n")
	lines = lines[:len(lines)-1]

	var result [][]string
	for _, v := range lines {
		line := strings.Split(v, " ")
		result = append(result, line)
	}

	return result
}

func findPos(tuples [][]string) int {
	var horizontal int
	var vertical int

	for _, tuple := range tuples {
		kind := tuple[0]
		num, _ := strconv.Atoi(tuple[1])

		switch kind {
		case "forward":
			horizontal += num
			break
		case "up":
			vertical -= num
			break
		case "down":
			vertical += num
			break
		}
	}

	return horizontal * vertical
}

func findPosWithAim(tuples [][]string) int {
	var horizontal int
	var vertical int
	var aim int

	for _, tuple := range tuples {
		kind := tuple[0]
		num, _ := strconv.Atoi(tuple[1])

		switch kind {
		case "forward":
			horizontal += num
			vertical += num * aim
			break
		case "up":
			aim -= num
			break
		case "down":
			aim += num
			break
		}
	}

	return horizontal * vertical
}
