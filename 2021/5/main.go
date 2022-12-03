package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

type line struct {
	start pos
	end   pos
}

type board struct {
	lines []line
	board [][]int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// eg
	// x1, y1      x2, y2
	// 105,697 -> 287,697
	// 705,62 -> 517,250

	scanner := bufio.NewScanner(f)

	var lines []line
	for scanner.Scan() {
		text := scanner.Text()

		coords := strings.Split(text, " -> ")
		start := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")

		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])
		endX, _ := strconv.Atoi(end[0])
		endY, _ := strconv.Atoi(end[1])

		lines = append(lines, line{
			start: pos{x: startX, y: startY},
			end:   pos{x: endX, y: endY},
		})
	}

	b := NewBoard(lines)
	b.drawHorizontal()
	b.drawVertical()
	fmt.Println(b.count())
	b.drawDiagonal()
	fmt.Println(b.count())

}

func (b board) count() int {
	count := 0
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			if b.board[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func (b *board) drawDiagonal() {
	// Firstly modify all lines so that they are always increasing in y
	dia := b.getLines(diagonal)
	for i := 0; i < len(dia); i++ {
		if dia[i].end.y < dia[i].start.y {
			dia[i].start, dia[i].end = dia[i].end, dia[i].start
		}
	}

	left, right := []line{}, []line{}
	for _, l := range dia {
		if l.start.x > l.end.x {
			left = append(left, l)
		} else {
			right = append(right, l)
		}
	}

	for _, l := range left {
		// draw leftwards diagonal
		x, y := l.start.x, l.start.y
		for x >= l.end.x {
			b.board[x][y]++
			x--
			y++
		}
	}

	for _, l := range right {
		// draw rightwards diagonal
		x, y := l.start.x, l.start.y
		for x <= l.end.x {
			b.board[x][y]++
			x++
			y++
		}
	}
}

func (b *board) drawHorizontal() {
	hoz := b.getLines(horizontal)

	// Make sure we're always iterating in the same direction i.e. ascending
	for i := 0; i < len(hoz); i++ {
		if hoz[i].end.x < hoz[i].start.x {
			hoz[i].start, hoz[i].end = hoz[i].end, hoz[i].start
		}
	}

	for _, l := range hoz {
		fmt.Printf("%d, %d -> %d, %d\n", l.start.x, l.start.y, l.end.x, l.end.y)
		x, y := l.start.x, l.start.y
		for x <= l.end.x {
			// Increment the value in the board
			b.board[x][y]++
			x++
		}
	}
}

func (b *board) drawVertical() {
	ver := b.getLines(vertical)

	// Make sure we're always iterating in the same direction i.e. ascending
	for i := 0; i < len(ver); i++ {
		if ver[i].end.y < ver[i].start.y {
			ver[i].start, ver[i].end = ver[i].end, ver[i].start
		}
	}

	for _, l := range ver {
		fmt.Printf("%d, %d -> %d, %d\n", l.start.x, l.start.y, l.end.x, l.end.y)
		x, y := l.start.x, l.start.y
		for y <= l.end.y {
			// Increment the value in the board
			b.board[x][y]++
			y++
		}
	}
}

func diagonal(l line) bool   { return l.start.x != l.end.x && l.start.y != l.end.y }
func horizontal(l line) bool { return l.start.y == l.end.y }
func vertical(l line) bool   { return l.start.x == l.end.x }

func (b board) getLines(pick func(line) bool) []line {
	var out []line
	for _, l := range b.lines {
		if pick(l) {
			out = append(out, l)
		}
	}
	return out
}

func NewBoard(lines []line) board {
	maxX, maxY := 0, 0

	for _, line := range lines {
		if line.start.x > maxX {
			maxX = line.start.x
		}
		if line.start.y > maxY {
			maxY = line.start.y
		}
		if line.end.x > maxX {
			maxX = line.end.x
		}
		if line.end.y > maxY {
			maxY = line.end.y
		}
	}

	grid := make([][]int, maxX+1)
	for i := 0; i <= maxX; i++ {
		grid[i] = make([]int, maxY+1)
	}

	fmt.Printf("Creating board %d x %d\n", maxX, maxY)

	return board{
		lines: lines,
		board: grid,
	}
}
