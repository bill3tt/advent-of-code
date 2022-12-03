package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type segment struct {
	input  []string
	output []string
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	segments := []segment{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		input := strings.Fields(line[0])
		for i := 0; i < len(input); i++ {
			input[i] = SortString(input[i])
		}
		output := strings.Fields(line[1])
		for i := 0; i < len(output); i++ {
			output[i] = SortString(output[i])
		}
		segments = append(segments, segment{input: input, output: output})
	}

	// In output, count how many times digits 1, 4, 7, 8 appear
	// 1 -> len 2
	// 4 -> len 4
	// 7 -> len 3
	// 8 -> len 7

	count := 0
	for _, s := range segments {
		for _, v := range s.output {
			switch len(v) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	sum := 0
	for _, s := range segments {
		mapping := solve(s)
		sum += decode(mapping, s.output)
	}
	fmt.Println(sum)
}

func decode(mapping map[string]int, encoded []string) int {
	var out []string
	for _, s := range encoded {
		out = append(out, fmt.Sprint(mapping[s]))
	}
	i, _ := strconv.Atoi(strings.Join(out, ""))
	fmt.Println(i)
	return i
}

func solve(seg segment) map[string]int {
	// Maintain bi-directional mapping between strings & ints
	stoi := map[string]int{}
	itos := map[int]string{}

	// Firstly, parse the easy ones
	for _, s := range seg.input {
		switch len(s) {
		case 2:
			stoi[s] = 1
			itos[1] = s
		case 3:
			stoi[s] = 7
			itos[7] = s
		case 4:
			stoi[s] = 4
			itos[4] = s
		case 7:
			stoi[s] = 8
			itos[8] = s
		}
	}

	// Next deal with the inputs of length 5
	var fivers []string
	for _, s := range seg.input {
		if len(s) == 5 {
			fivers = append(fivers, s)
		}
	}

	for _, f := range fivers {
		four, ok := itos[4]
		if !ok {
			log.Fatalf("Could not find string for integer 4")
		}
		if common(f, four) == 2 {
			stoi[f] = 2
			itos[2] = f
		}
	}

	for _, f := range fivers {
		two, ok := itos[2]
		if !ok {
			log.Fatalf("Could not find string for integer 2")
		}
		switch common(f, two) {
		case 4:
			stoi[f] = 3
			itos[3] = f
		case 3:
			stoi[f] = 5
			itos[5] = f
		}
	}

	var sixers []string
	for _, s := range seg.input {
		if len(s) == 6 {
			sixers = append(sixers, s)
		}
	}

	for _, s := range sixers {
		seven := itos[7]
		four := itos[4]
		five, ok := itos[5]
		if !ok {
			log.Fatalf("Do not have string for 5")
		}

		if common(s, seven) == 2 {
			itos[6] = s
			stoi[s] = 6
		}
		if common(s, four) == 4 {
			itos[9] = s
			stoi[s] = 9
		}
		if common(s, five) == 4 {
			itos[0] = s
			stoi[s] = 0
		}
	}

	return stoi
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// common returns the number of characters in common between s & c
func common(s, t string) int {
	present := map[byte]bool{}
	for i := range s {
		present[s[i]] = true
	}
	total := 0
	for i := range t {
		if present[t[i]] {
			total++
		}
	}
	return total
}
