package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var input []int
	for _, s := range strings.Split(scanner.Text(), ",") {
		f, _ := strconv.Atoi(s)
		input = append(input, f)
	}
	fmt.Println(len(input))

	// Firstly we need to convert the input array into an actual position count array.
	// What is the max value in the input?
	max := 0
	for _, i := range input {
		if i > max {
			max = i
		}
	}

	// Make the pos slice one bigger than max so we can index max
	pos := make([]int, max+1)
	fmt.Println(max)

	// Instantiate our pos slice
	for _, i := range input {
		pos[i]++
	}

	left := make([]int, len(pos))

	// Calculate total energy needed to move to given position from left
	for i := 0; i < len(pos); i++ {
		cost := 1
		for j := i + 1; j < len(pos); j++ {
			left[j] += (pos[i] * (j - i)) * cost
			cost++
		}
	}

	right := make([]int, len(pos))
	// Calculate total energy needed to move to given position from right
	for i := len(pos) - 1; i >= 0; i-- {
		cost := 1
		for j := i - 1; j >= 0; j-- {
			right[j] += (pos[i] * (i - j)) * cost
			cost++
		}
	}

	min := math.MaxInt
	var minPos int
	for i := 0; i < len(pos); i++ {
		total := left[i] + right[i]
		if total < min {
			min = total
			minPos = i
		}
	}
	fmt.Println(minPos, min)
}
