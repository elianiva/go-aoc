package day10

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func Solve() {
	fmt.Println("======== DAY 10 ========")

	inputFile, err := ioutil.ReadFile("day10/input.txt")
	if err != nil {
		panic(err)
	}

	input := strToLines(inputFile)

	errorScore := findErrorScore(input)
	completionScore := findCompletionScore(input)

	fmt.Println("PART 1:", errorScore)
	fmt.Println("PART 2:", completionScore)
}

func strToLines(str []byte) []string {
	lines := strings.Split(string(str), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

func findErrorScore(lines []string) int {
	var totalScore int
	for _, line := range lines {
		totalScore += getLineErrorScore(line)
	}
	return totalScore
}

func findCompletionScore(lines []string) int {
	var totalScore sort.IntSlice
	for _, line := range lines {
		if getLineErrorScore(line) == 0 {
			totalScore = append(totalScore, getLineCompletionScore(line))
		}
	}
	totalScore.Sort()

	length := totalScore.Len()
	arithMean := (length / 2) + (length % 2)

	return totalScore[arithMean-1]
}

// Stack is a stack data structure for `byte`
type Stack struct {
	items []byte
}

func (s *Stack) Push(c byte) {
	s.items = append(s.items, c)
}

func (s *Stack) Pop() byte {
	if len(s.items) > 0 {
		popped := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return popped
	}
	return 0
}

func (s *Stack) Peek() byte {
	return s.items[len(s.items)-1]
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Get() []byte {
	return s.items
}

func isOpening(c byte) bool {
	return c == '(' || c == '[' || c == '<' || c == '{'
}

func isClosing(c byte) bool {
	return c == ')' || c == ']' || c == '>' || c == '}'
}

func getPair(c byte) byte {
	return map[byte]byte{
		')': '(', '(': ')',
		']': '[', '[': ']',
		'}': '{', '{': '}',
		'>': '<', '<': '>',
	}[c]
}

func getCharErrorScore(c byte) int {
	return map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}[c]
}

func getLineErrorScore(line string) int {
	stack := &Stack{}

	var score int
	for _, c := range line {
		char := byte(c)
		if isOpening(char) {
			stack.Push(char)
		} else if isClosing(char) {
			if getPair(char) == stack.Peek() {
				stack.Pop()
			} else {
				// idk this message just looks cool
				if os.Getenv("DEBUG") == "1" {
					fmt.Printf("Expected %s, but found %s instead.\n", string(getPair(stack.Peek())), string(char))
				}
				score += getCharErrorScore(char)
				break
			}
		}
	}

	return score
}

func getCharCompletionScore(c byte) int {
	return map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}[c]
}

func getLineCompletionScore(line string) int {
	stack := &Stack{}

	for _, c := range line {
		char := byte(c)
		if isOpening(char) {
			stack.Push(char)
		} else if isClosing(char) {
			if getPair(char) == stack.Peek() {
				stack.Pop()
			}
		}
	}

	var score int
	chars := stack.Get()
	for i := len(chars) - 1; i >= 0; i-- {
		score = (score * 5) + getCharCompletionScore(getPair(chars[i]))
	}

	return score
}
