package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type bingo struct {
	picks  []int
	boards []board
}

type board [][]tile

type tile struct {
	val   int
	drawn bool
}

func main() {
	// Firstly read the input.
	// Format is as follows:
	// * list of comma separated numbers that are the choices
	// * newline delimited boards

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	b := parseBingo(f)
	fmt.Println(b.picks)
	b.play()
}

func (b *bingo) play() {
	for _, p := range b.picks {
		fmt.Printf("Pick: %d\n", p)
		// Mark picks in boards as drawn
		for i := 0; i < len(b.boards); i++ {
			b.boards[i].mark(p)
		}
		// Check for any winners
		for i := 0; i < len(b.boards); i++ {
			if b.boards[i].wins() {
				// Calculate the final score & print
				fmt.Printf("Board %d wins.\n", i)
				fmt.Println(b.boards[i].score(p))
				fmt.Println(b.boards[i])
				return
			}
		}
	}
}

func (b *board) score(pick int) int {
	sum := 0
	for i := 0; i < len(*b); i++ {
		for j := 0; j < len((*b)[0]); j++ {
			if !(*b)[i][j].drawn {
				sum += (*b)[i][j].val
			}
		}
	}
	return pick * sum
}

func (b *board) wins() bool {
	// Check rows for all drawn
RowLoop:
	for i := 0; i < len(*b); i++ {
		for j := 0; j < len((*b)[i]); j++ {
			if !(*b)[i][j].drawn {
				continue RowLoop
			}
		}
		return true
	}
	// Check colums for all drawn
ColumnLoop:
	for i := 0; i < len((*b)[0]); i++ {
		for j := 0; j < len(*b); j++ {
			if !(*b)[j][i].drawn {
				continue ColumnLoop
			}
		}
		return true
	}
	return false
}

func (b *board) mark(pick int) {
	for i := 0; i < len(*b); i++ {
		for j := 0; j < len((*b)[0]); j++ {
			if (*b)[i][j].val == pick {
				(*b)[i][j].drawn = true
			}
		}
	}
}

func (b board) String() string {
	out := ""
	for i := 0; i < len(b); i++ {
		line := ""
		for j := 0; j < len((b)[i]); j++ {
			s := fmt.Sprintf("%4d", (b)[i][j].val)
			if (b)[i][j].drawn {
				line += color.GreenString(s)
			} else {
				line += s
			}
		}
		line += "\n"
		out += line
	}
	return out
}

func parseBingo(f *os.File) bingo {

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	picks := []int{}
	for _, s := range strings.Split(scanner.Text(), ",") {
		p, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		picks = append(picks, p)
	}
	fmt.Printf("Got %d picks\n", len(picks))
	fmt.Println(picks)

	// Skip the newlines before the board data
	scanner.Scan()

	// boards := []board{}
	// Read the boards into memory
	// They are in 5x5 grid format
	boards := []board{board{}}
	for scanner.Scan() {
		lineString := scanner.Text()
		if len(lineString) == 0 {
			boards = append(boards, board{})
			continue
		}
		var line []tile
		tokens := strings.Fields(scanner.Text())
		for _, s := range tokens {
			v, _ := strconv.Atoi(s)
			line = append(line, tile{val: v, drawn: false})
		}
		boards[len(boards)-1] = append(boards[len(boards)-1], line)
	}

	fmt.Printf("Got %d boards.\n", len(boards))

	for _, b := range boards {
		fmt.Println(b)
	}

	return bingo{
		picks:  picks,
		boards: boards,
	}
}
