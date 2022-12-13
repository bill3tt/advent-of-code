package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	direction Direction
	distance  int
}

type Pos struct {
	x, y int
}

type Direction string

const (
	U Direction = "U"
	D Direction = "D"
	L Direction = "L"
	R Direction = "R"
)

func (p *Pos) Move(d Direction) {
	switch d {
	case U:
		p.y++
	case D:
		p.y--
	case L:
		p.x--
	case R:
		p.x++
	}
}

func (t *Pos) Follow(h Pos) {
	// Are the two positions 'touching' i.e. adjacent? Do nothing
	if touching(*t, h) {
		return
	}

	// Are the two positions in the same row or column? Move towards.
	if t.x == h.x && t.y < h.y {
		t.y++
	}
	if t.x == h.x && t.y > h.y {
		t.y--
	}

	if t.y == h.y && t.x < h.x {
		t.x++
	}
	if t.y == h.y && t.x > h.x {
		t.x--
	}

	// Else they are not touching and not in the same row or column. Move diagonally
	if t.x < h.x && t.y < h.y {
		t.x++
		t.y++
	}
	if t.x < h.x && t.y > h.y {
		t.x++
		t.y--
	}
	if t.x > h.x && t.y < h.y {
		t.x--
		t.y++
	}
	if t.x > h.x && t.y > h.y {
		t.x--
		t.y--
	}
}

func touching(a, b Pos) bool {
	return abs(a.x-b.x) <= 1 && abs(a.y-b.y) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var commands []Command

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		distance, _ := strconv.Atoi(parts[1])
		commands = append(commands, Command{direction: Direction(parts[0]), distance: distance})
	}
	fmt.Printf("Part 1: Visited %d spaces.\n", simulate(commands, 2))
	fmt.Printf("Part 2: Visited %d spaces.\n", simulate(commands, 10))
}

func simulate(commands []Command, length int) int {
	var rope []*Pos
	for i := 0; i < length; i++ {
		rope = append(rope, &Pos{0, 0})
	}

	counter := map[Pos]int{}
	for _, c := range commands {
		for d := c.distance; d > 0; d-- {
			head, tail := rope[0], rope[length-1]
			head.Move(c.direction)
			for i := 1; i < length; i++ {
				rope[i].Follow(*rope[i-1])
			}
			counter[Pos{tail.x, tail.y}]++
		}
	}
	return len(counter)
}
