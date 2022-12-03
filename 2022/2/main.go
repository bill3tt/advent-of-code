package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	var rounds []string
	for scanner.Scan() {
		rounds = append(rounds, scanner.Text())
	}

	fmt.Printf("Part 1 score: %d\n", calculate(rounds))
	converted := part2(rounds)
	fmt.Printf("Part 2 score: %d\n", calculate(converted))
}

func part2(rounds []string) []string {
	// X means you need to loos	// Y means you need to draw
	// Z means you need to win

	// Modify the input round slice to correspond with the desired outcome
	// Convert the input to the same format used in part 1 then reuse the calucation mechanism.

	var converted []string

	for _, round := range rounds {
		switch round {
		// They have Rock, I need to Loose, I need Scissors
		case "A X":
			converted = append(converted, "A Z")
		// They have Rock, I need to Draw, I need Rock
		case "A Y":
			converted = append(converted, "A X")
		// They have Rock, I need to Win, I need Paper
		case "A Z":
			converted = append(converted, "A Y")
		// They have Paper, I need to Loose, I need Rock
		case "B X":
			converted = append(converted, "B X")
		// They have Paper, I need to Draw, I need Paper
		case "B Y":
			converted = append(converted, "B Y")
		// They have Paper, I need to Win, I need Scissors
		case "B Z":
			converted = append(converted, "B Z")
		// They have Scissors, I need to Loose, I need Paper
		case "C X":
			converted = append(converted, "C Y")
		// They have Scissors, I need to Draw, I need Scissors
		case "C Y":
			converted = append(converted, "C Z")
		// They have Scissors, I need to Win, I need Rock
		case "C Z":
			converted = append(converted, "C X")
		default:
			panic("unexpected input")
		}
	}
	return converted
}

func calculate(rounds []string) int {
	score := 0
	for _, round := range rounds {
		switch round {
		// Rock versus Rock
		case "A X":
			score += (1 + 3)
		// Rock versus Paper
		case "A Y":
			score += (2 + 6)
		// Rock versus Scissors
		case "A Z":
			score += (3 + 0)
		// Paper versus Rock
		case "B X":
			score += (1 + 0)
		// Paper versus Paper
		case "B Y":
			score += (2 + 3)
		// Paper versus Scissors
		case "B Z":
			score += (3 + 6)
		// Scisscors versus Rock
		case "C X":
			score += (1 + 6)
		// Scisscors versus Paper
		case "C Y":
			score += (2 + 0)
		// Scisscors versus Scissors
		case "C Z":
			score += (3 + 3)
		default:
			panic("unexpected input")
		}
	}
	return score
}
