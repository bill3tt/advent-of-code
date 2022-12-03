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

	fmt.Println(groups[0])

}

func part1(rucksacks []string) {
	total := 0

	for _, r := range rucksacks {
		left, right := r[:len(r)/2], r[len(r)/2:]

		// Find the characters that appear in both compartments
		seen := map[byte]bool{}
		for i := range left {
			seen[left[i]] = true
		}

		var common byte
		for i := range right {
			if _, ok := seen[right[i]]; ok {
				common = right[i]
			}
		}
		total += priority(common)
	}

	fmt.Printf("Part 1: %d\n", total)
}

func priority(b byte) int {
	if int(b) >= 97 {
		return int(b) - int(byte('a')) + 1
	}
	return int(b) - int(byte('A')) + 27
}
