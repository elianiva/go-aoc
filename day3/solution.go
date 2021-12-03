package day3

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("======== DAY 3 ========")

	inputFile, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(inputFile), "\n")
	input = input[:len(input)-1]

	powerConsumption := findPowerConsumption(input)
	lifeSupport := findLifeSupport(input)

	fmt.Println("PART 1:", powerConsumption)
	fmt.Println("PART 2:", lifeSupport)
}

// will return negative or positive integer depending on which bit occurs the most
// positive int means 1, negative int means 0
func findMostOccurences(lines []string) []int {
	result := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, char := range line {
			num, _ := strconv.Atoi(string(char))
			if num > 0 {
				result[i] += 1
			} else {
				result[i] -= 1
			}
		}
	}

	return result
}

func findPowerConsumption(lines []string) int {
	occurences := findMostOccurences(lines)

	var g strings.Builder
	var e strings.Builder
	for _, v := range occurences {
		if v > 0 {
			g.WriteString("1")
			e.WriteString("0")
		} else {
			g.WriteString("0")
			e.WriteString("1")
		}
	}

	gamma, _ := strconv.ParseInt(g.String(), 2, 0)
	epsilon, _ := strconv.ParseInt(e.String(), 2, 0)

	return int(gamma * epsilon)
}

func findLastBit(lines []string, pos int, bit []string) int {
	if len(lines) == 1 {
		result, _ := strconv.ParseInt(lines[0], 2, 0)
		return int(result)
	}

	var ones int
	var zeroes int
	for _, line := range lines {
		if string(line[pos]) == "1" {
			ones += 1
		} else {
			zeroes += 1
		}
	}

	var filteredLines []string
	// not exactly proud of this thing
	for _, line := range lines {
		if ones > zeroes {
			if string(line[pos]) == bit[0] {
				filteredLines = append(filteredLines, line)
			}
			if len(filteredLines) == ones {
				break
			}
		} else if zeroes > ones {
			if string(line[pos]) == bit[1] {
				filteredLines = append(filteredLines, line)
			}
			if len(filteredLines) == zeroes {
				break
			}
		} else {
			if string(line[pos]) == bit[0] {
				filteredLines = append(filteredLines, line)
				break
			}
		}
	}
	return findLastBit(filteredLines, pos+1, bit)
}

func findLifeSupport(lines []string) int {
	oxygen := findLastBit(lines, 0, []string{"1", "0"})
	scrubber := findLastBit(lines, 0, []string{"0", "1"})

	return oxygen * scrubber
}
