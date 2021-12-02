package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("======== DAY 1 ========")

	inputFile, err := ioutil.ReadFile("day1/input1.txt")
	if err != nil {
		panic(err)
	}

	input := linesToSlice(inputFile)

	depth := findDepth(input)
	depthChunk := findDepthChunk(input)

	fmt.Println("PART 1:", depth)
	fmt.Println("PART 2:", depthChunk)
}

func linesToSlice(str []byte) []int {
	lines := strings.Split(string(str), "\n")
	lines = lines[:len(lines)-1]

	var result []int
	for _, v := range lines {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}

	return result
}

func findDepth(lines []int) int {
	var result int

	highest := lines[0]
	for i := 1; i < len(lines); i++ {
		if lines[i] > highest {
			result += 1
		}

		highest = lines[i]
	}

	return result
}

func findDepthChunk(lines []int) int {
	var sums []int

	// we want it to stop after it reaches the last 3 elements
	// suppose we have this -> [1, 2, 3, 4, 5, 6]
	// we want it to be [1,2,3] - [2,3,4] - [3,4,5] - [4,5,6]
	// hence we need the -2 to remove 5 and 6 from the array.
	for i := 0; i < len(lines)-2; i++ {
		// this thing is just for the sake of readability
		// gofmt is forcing me to do lines[i]+lines[i+1]+lines... which isn't easy to read
		a := lines[i]
		b := lines[i+1]
		c := lines[i+2]
		sums = append(sums, a+b+c)
	}

	return findDepth(sums)
}
