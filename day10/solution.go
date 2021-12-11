package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Solve() {
	fmt.Println("======== DAY 10 ========")

	inputFile, err := os.Open("day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	errorScore, completionScore := findScore(scanner)

	fmt.Println("PART 1:", errorScore)
	fmt.Println("PART 2:", completionScore)
}

func strToLines(str []byte) []string {
	lines := strings.Split(string(str), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

func findScore(scanner *bufio.Scanner) (int, int) {
	var completionTotal sort.IntSlice
	var errorTotal int
	for scanner.Scan() {
		line := scanner.Text()
		errScore := getLineErrorScore(line)
		if errScore == 0 {
			completionTotal = append(completionTotal, getLineCompletionScore(line))
		} else {
			errorTotal += getLineErrorScore(line)
		}
	}
	completionTotal.Sort()

	length := completionTotal.Len()
	arithMean := (length / 2) + (length % 2)

	return errorTotal, completionTotal[arithMean-1]
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
	// it was using map[byte]byte, but switch statement is more efficient
	// thanks @smolck
	switch c {
	case '(':
		return ')'
	case ')':
		return '('
	case '[':
		return ']'
	case ']':
		return '['
	case '{':
		return '}'
	case '}':
		return '{'
	case '>':
		return '<'
	case '<':
		return '>'
	default:
		return ' '
	}
}

func getCharErrorScore(c byte) int {
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return 0
	}
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
	switch c {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	default:
		return 0
	}
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
		score = (score * 5) + getCharCompletionScore(chars[i])
	}

	return score
}
