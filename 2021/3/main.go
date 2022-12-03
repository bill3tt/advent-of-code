package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	oxygen := make([]string, len(input))
	copy(oxygen, input)

	for i := 0; i < len(oxygen[0]); i++ {
		c := mostCommon(oxygen, i)
		switch c {
		case "1":
			oxygen = filter(oxygen, i, "1")
		case "0":
			oxygen = filter(oxygen, i, "0")
		}

	}

	fmt.Println(len(oxygen))

	scrubber := make([]string, len(input))
	copy(scrubber, input)

	for i := 0; i < len(scrubber[0]); i++ {
		c := mostCommon(scrubber, i)
		switch c {
		case "1":
			scrubber = filter(scrubber, i, "0")
		case "0":
			scrubber = filter(scrubber, i, "1")
		}
		if len(scrubber) == 1 {
			break
		}
	}

	g, err := strconv.ParseInt(oxygen[0], 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(scrubber[0], 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(g * e)
}

func filter(input []string, index int, char string) []string {
	var out []string
	for _, s := range input {
		if string(s[index]) == char {
			out = append(out, s)
		}
	}
	return out
}

func mostCommon(input []string, index int) string {
	count := map[string]int{}
	for j := 0; j < len(input); j++ {
		count[string(input[j][index])]++
	}
	var mostCommon string
	var largest int
	for k, v := range count {
		if v > largest {
			mostCommon = k
			largest = v
		}
	}
	return mostCommon
}
