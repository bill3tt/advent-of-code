package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var rucksacks []string
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}

	part1(rucksacks)
	part2(rucksacks)
}

func part2(rucksacks []string) {
	// Firstly group rucksacks into groups
	groups := [][]string{}
	for i := 0; i < len(rucksacks)/3; i++ {
		groups = append(groups, []string{})
	}

	for i, r := range rucksacks {
		groups[i/3] = append(groups[i/3], r)
	}

	total := 0
	for _, g := range groups {
		total += priority(findCommonChar(g))
	}
	fmt.Printf("Part 1: %d\n", total)
}

func part1(rucksacks []string) {
	total := 0

	for _, r := range rucksacks {
		left, right := r[:len(r)/2], r[len(r)/2:]
		total += priority(findCommonChar([]string{left, right}))
	}

	fmt.Printf("Part 1: %d\n", total)
}

// findCommonChar will find the single byte that is common across the input strings
// The puzzle input is guaranteed to have only one common character
func findCommonChar(inputs []string) byte {
	common := map[byte]struct{}{}

	// create the first hashmap
	for i := range inputs[0] {
		common[inputs[0][i]] = struct{}{}
	}

	var next map[byte]struct{}

	for i := 1; i < len(inputs); i++ {
		// Initialise the next seen map to be empty
		next = map[byte]struct{}{}
		for j, _ := range inputs[i] {
			b := inputs[i][j]
			// If this character is in the previous map, its common to the previous inputs
			if _, ok := common[b]; ok {
				next[b] = struct{}{}
			}
		}
		common = next
	}
	out := []byte{}
	for k := range common {
		out = append(out, k)
	}
	return out[0]
}

func priority(b byte) int {
	if int(b) >= 97 {
		return int(b) - int(byte('a')) + 1
	}
	return int(b) - int(byte('A')) + 27
}
