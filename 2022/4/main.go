package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var assignments [][]int
	for scanner.Scan() {
		var line []int
		parts := strings.FieldsFunc(scanner.Text(), splitter)
		for _, p := range parts {
			i, _ := strconv.Atoi(p)
			line = append(line, i)
		}
		assignments = append(assignments, line)
	}

	enclosed := 0
	overlapping := 0
	for _, a := range assignments {
		aStart, aEnd, bStart, bEnd := a[0], a[1], a[2], a[3]

		// Does one range fully enclose the other?
		if (aStart <= bStart && aEnd >= bEnd) || (bStart <= aStart && bEnd >= aEnd) {
			enclosed++
		}

		// Do the ranges overlap at all?
		if (aEnd >= bStart && bStart >= aStart) || (bEnd >= aStart && aStart >= bStart) {
			overlapping++
		}
	}
	fmt.Printf("Part 1: %d\n", enclosed)
	fmt.Printf("Part 2: %d\n", overlapping)
}

func splitter(r rune) bool {
	switch r {
	case '-':
		return true
	case ',':
		return true
	default:
		return false
	}
}
