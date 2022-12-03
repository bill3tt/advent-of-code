package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct{ x, y int }

type queue []pos

func (q *queue) push(p pos) {
	*q = append(*q, p)
}

func (q *queue) pop() pos {
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}

func main() {
	grid := [][]int{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		row := []int{}
		for _, s := range strings.Split(scanner.Text(), "") {
			val, _ := strconv.Atoi(s)
			row = append(row, val)
		}
		grid = append(grid, row)
	}

	fmt.Println(grid)

	// First, energy level of each octopus increases by 1
	// Any octopus with energy > 9 flashes.
	//	* This increases energy of all adjacent octopuses by 1 including diagonally adjacent
	//	* This process continues as long as new octopuses keep having their energy level above 9
	// Any octopus that flashed has its energy set to 0

	yMax, xMax := len(grid), len(grid[0])
	totalFlashes := 0

	step := 1
	for step <= 1000 {
		octopuses := queue{}
		stepFlashes := 0

		// First step add all of the grid locations to the queue
		for y := 0; y < yMax; y++ {
			for x := 0; x < xMax; x++ {
				octopuses.push(pos{y: y, x: x})
			}
		}
		// Process all the elements in the flash queue.
		flashed := map[pos]bool{}

		for len(octopuses) > 0 {
			o := octopuses.pop()

			// Firstly check if this is a valid location
			if o.y >= yMax || o.x >= xMax || o.y < 0 || o.x < 0 {
				continue
			}

			// Increment this octopuses energy by 1
			grid[o.y][o.x]++

			if grid[o.y][o.x] <= 9 {
				continue
			}

			// If this position has already flashed, continue
			if _, ok := flashed[o]; ok {
				continue
			}

			// Increment all adjacent locations & add them to the flash queue if need be
			octopuses.push(pos{y: o.y - 1, x: o.x - 1})
			octopuses.push(pos{y: o.y - 1, x: o.x})
			octopuses.push(pos{y: o.y - 1, x: o.x + 1})
			octopuses.push(pos{y: o.y, x: o.x - 1})
			octopuses.push(pos{y: o.y, x: o.x + 1})
			octopuses.push(pos{y: o.y + 1, x: o.x - 1})
			octopuses.push(pos{y: o.y + 1, x: o.x})
			octopuses.push(pos{y: o.y + 1, x: o.x + 1})

			flashed[o] = true
			totalFlashes++
			stepFlashes++
		}

		if stepFlashes == 100 {
			fmt.Printf("Step: %d, Flashes %d\n", step, stepFlashes)
			break
		}

		// Reset all positions that have flashed to 0
		for y := 0; y < yMax; y++ {
			for x := 0; x < xMax; x++ {
				if grid[y][x] > 9 {
					grid[y][x] = 0
				}
			}
		}

		step++
	}

	// Find the number of flashes after 100 steps
	fmt.Println(totalFlashes)
}
