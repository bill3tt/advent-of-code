package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	distance  int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var commands []command
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		direction := input[0]
		distance, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, command{
			direction, distance,
		})
	}

	horizontal, depth, aim := 0, 0, 0

	for _, c := range commands {
		switch c.direction {
		case "forward":
			horizontal += c.distance
			depth += (aim * c.distance)
		case "down":
			aim += c.distance
		case "up":
			aim -= c.distance
		}
	}
	fmt.Println(horizontal * depth)
}
