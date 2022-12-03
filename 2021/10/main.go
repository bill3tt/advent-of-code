package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack []string

func (s *stack) push(b string) {
	*s = append(*s, b)
}

func (s *stack) pop() string {
	p := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return p
}

var closer = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
	"<": ">",
}

var scoring = map[string]int{
	"}": 1197,
	"]": 57,
	")": 3,
	">": 25137,
}

func main() {

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var input [][]string
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}

	score := 0

LineLoop:
	for _, line := range input {
		s := stack{}
		for _, i := range line {
			if isOpener(i) {
				s.push(i)
			} else {
				// We have a closer
				opener := s.pop()
				if !match(opener, i) {
					score += scoring[i]
					continue LineLoop
				}
			}
		}
	}

	fmt.Println(score)

	// Corrupted - wrong closer

	// Scoring
	// ): 3 points.
	// ]: 57 points.
	// }: 1197 points.
	// >: 25137 points.

}

func isOpener(s string) bool {
	if _, ok := closer[s]; !ok {
		return false
	}
	return true
}

func match(s, t string) bool {
	if v := closer[s]; v == t {
		return true
	}
	return false
}
