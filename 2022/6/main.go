package main

import (
	"bufio"
	"fmt"
	"os"
)

const PACKET_MARKER_WIDTH = 4
const MESSAGE_MARKER_WIDTH = 14

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}

	// In the protocol being used by the Elves, the start of a packet is indicated by a sequence of four characters that are all different.
	fmt.Printf("Input is %d characters long.\n", len(input))

	// Part 1
	findMarker(input, PACKET_MARKER_WIDTH)

	// Part 2
	findMarker(input, MESSAGE_MARKER_WIDTH)
}

func findMarker(input string, headerLength int) {
	count := map[byte]int{}

	var i int

	// Pre-populate the fields
	for i := 0; i < headerLength; i++ {
		count[input[i]]++
	}

	for i = headerLength; i < len(input); i++ {
		if noDuplicates(count) {
			fmt.Printf("First header of length %d is %s, idenitified after processing %d characters.\n", headerLength, input[i-4:i], i)
			break
		}

		count[input[i]]++
		count[input[i-headerLength]]--
	}
}

func noDuplicates(counter map[byte]int) bool {
	for _, v := range counter {
		if v > 1 {
			return false
		}
	}
	return true
}
