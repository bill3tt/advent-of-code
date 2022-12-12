package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Summary struct {
	max, visible int
}

type Pos struct {
	x, y, h               int
	left, right, up, down Summary
	trees                 *Trees
}

type Trees struct {
	grid          [][]*Pos
	height, width int
}

func (p Pos) visible() bool {
	return p.h > p.left.max ||
		p.h > p.up.max ||
		p.h > p.down.max ||
		p.h > p.right.max ||
		p.x == 0 ||
		p.y == 0 ||
		p.x == p.trees.width-1 ||
		p.y == p.trees.height-1
}

func (p Pos) scenicScore() int {
	return p.left.visible * p.up.visible * p.down.visible * p.right.visible
}

func (p Pos) String() string {
	var v string
	if p.visible() {
		v = "v"
	} else {
		v = " "
	}
	return fmt.Sprintf(`
	    ↑%d,%d   
←%d,%d	 %d%s (%d,%d) %d,%d→
	    ↓%d,%d
	`, p.up.max, p.up.visible, p.left.max, p.left.visible, p.h, v, p.x, p.y, p.right.max, p.right.visible, p.down.max, p.down.visible)
}

func (t *Trees) valid(x, y int) bool {
	return x >= 0 && y >= 0 && x < t.width && y < t.height
}

func (t *Trees) findMaxHeight(p *Pos, iterator func(start *Pos) *Pos) int {
	next := iterator(p)
	highest := 0
	for next != nil {
		if next.h > highest {
			highest = next.h
		}
		next = iterator(next)
	}
	return highest
}

// Iterator function to return the position above or nil if invalid
func (t *Trees) down(p *Pos) *Pos {
	if t.valid(p.x, p.y+1) {
		return t.grid[p.y+1][p.x]
	}
	return nil
}

// Iterator function to return the position above or nil if invalid
func (t *Trees) up(p *Pos) *Pos {
	if t.valid(p.x, p.y-1) {
		return t.grid[p.y-1][p.x]
	}
	return nil
}

// Iterator function to return the position above or nil if invalid
func (t *Trees) left(p *Pos) *Pos {
	if t.valid(p.x-1, p.y) {
		return t.grid[p.y][p.x-1]
	}
	return nil
}

// Iterator function to return the position above or nil if invalid
func (t *Trees) right(p *Pos) *Pos {
	if t.valid(p.x+1, p.y) {
		return t.grid[p.y][p.x+1]
	}
	return nil
}

func (t *Trees) findVisible(start *Pos, iterator func(start *Pos) *Pos) int {
	next := iterator(start)
	visible := 0
	// If we are on the boundary
	if next == nil {
		return visible
	}
	for next != nil {
		visible++
		if next.h >= start.h {
			break
		}
		next = iterator(next)
	}
	return visible

}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	trees := &Trees{height: len(input), width: len(input[0]), grid: [][]*Pos{}}

	for y, line := range input {
		var row []*Pos
		for x, v := range line {
			h, _ := strconv.Atoi(string(v))
			row = append(row, &Pos{x: x, y: y, h: h, trees: trees})
		}
		trees.grid = append(trees.grid, row)
	}

	for y := 0; y < trees.height; y++ {
		for x := 0; x < trees.width; x++ {
			if x == 2 && y == 3 {
				fmt.Println("hiya")
			}
			p := trees.grid[y][x]
			p.up = Summary{
				max:     trees.findMaxHeight(p, trees.up),
				visible: trees.findVisible(p, trees.up),
			}
			p.down = Summary{
				max:     trees.findMaxHeight(p, trees.down),
				visible: trees.findVisible(p, trees.down),
			}
			p.left = Summary{
				max:     trees.findMaxHeight(p, trees.left),
				visible: trees.findVisible(p, trees.left),
			}
			p.right = Summary{
				max:     trees.findMaxHeight(p, trees.right),
				visible: trees.findVisible(p, trees.right),
			}
		}
	}

	visible := 0
	maxScenicScore := 0
	for y := 0; y < trees.height; y++ {
		for x := 0; x < trees.width; x++ {
			p := trees.grid[y][x]
			if p.visible() {
				visible++
			}
			if p.scenicScore() > maxScenicScore {
				maxScenicScore = p.scenicScore()
			}
		}
	}
	fmt.Printf("Part 1 - visible trees: %d\n", visible)
	fmt.Printf("Part 2 - max scenic score: %d\n", maxScenicScore)
}
