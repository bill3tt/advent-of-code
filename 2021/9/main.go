package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type cave struct {
	grid [][]pos
}

type pos struct {
	x, y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var grid [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := []int{}
		text := scanner.Text()
		for i := 0; i < len(text); i++ {
			c := string(text[i])
			d, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, d)
		}
		grid = append(grid, line)
	}

	riskLevel := 0
	// Find the low points, low points are lower than any of the 4 adjacent points
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if risky(grid, x, y) {
				riskLevel += (grid[y][x] + 1)
			}
		}
	}
	fmt.Println(riskLevel)
}

// [
//	[2 1 9 9 9 4 3 2 1 0]
//  [3 9 8 7 8 9 4 9 2 1]
//  [9 8 5 6 7 8 9 8 9 2]
// ]

// height = 3
// width = 10

func risky(grid [][]int, x, y int) bool {
	var adjacent []int
	height, width := len(grid), len(grid[0])

	// Left
	if x > 0 {
		adjacent = append(adjacent, grid[y][x-1])
	}
	// Right
	if x < width-1 {
		adjacent = append(adjacent, grid[y][x+1])
	}
	// Up
	if y > 0 {
		adjacent = append(adjacent, grid[y-1][x])
	}
	// Down
	if y < height-1 {
		adjacent = append(adjacent, grid[y+1][x])
	}

	for _, a := range adjacent {
		// If the grid location is greater than any of the adjacent points, it is not risky
		if grid[y][x] >= a {
			return false
		}
	}
	fmt.Println(y, x)
	return true
}
