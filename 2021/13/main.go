package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type fold struct {
	direction string
	value     int
}

type pos struct {
	x, y int
}

func (p *pos) fold(f fold) {
	if f.direction == "x" && p.x > f.value {
		p.x = f.value - (p.x - f.value)
	}
	if f.direction == "y" && p.y > f.value {
		p.y = f.value - (p.x - f.value)
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var points []*pos
	var folds []fold

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// e.g. 1020,836
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			points = append(points, &pos{x, y})

		}

		// e.g. fold along y=223
		if strings.Contains(line, "fold along") {
			parts := strings.FieldsFunc(line, splitter)
			v, _ := strconv.Atoi(parts[3])
			folds = append(folds, fold{parts[2], v})
		}
	}

	for _, fold := range folds {
		for _, point := range points {
			point.fold(fold)
		}
	}
	plot(points)
}

// unique returns the total unique points
func unique(positions []*pos) int {
	counter := map[pos]int{}
	for _, p := range positions {
		counter[*p]++
	}
	fmt.Println(counter)
	return len(counter)
}

func splitter(r rune) bool {
	return r == ' ' || r == '='
}

func plot(positions []*pos) {
	// firstly find max x & y values
	xmax, ymax := 0, 0
	for _, p := range positions {
		if p.x > xmax {
			xmax = p.x
		}
		if p.y > ymax {
			ymax = p.y
		}
	}
	fmt.Println(xmax, ymax)

	var lines [][]byte
	for y := 0; y <= ymax; y++ {
		line := []byte{}
		for x := 0; x <= xmax; x++ {
			line = append(line, byte(' '))
		}
		lines = append(lines, line)
	}

	for _, p := range positions {
		lines[p.y][p.x] = '*'
	}

	for _, l := range lines {
		fmt.Println(string(l[:]))
	}
}
