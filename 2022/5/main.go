package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const NUM_STACKS = 9

type instruction struct {
	quantity, from, to int
}

func (i instruction) String() string {
	return fmt.Sprintf("From: %d, To: %d, Quantity: %d", i.from, i.to, i.quantity)
}

type stack []string

func (st *stack) headPush(s string) {
	(*st) = append([]string{s}, (*st)...)
}

func (st *stack) pop() string {
	ret := (*st)[len(*st)-1]
	(*st) = (*st)[:len(*st)-1]
	return ret
}

func (st *stack) push(d string) {
	(*st) = append((*st), d)
}

func (st *stack) pick(n int) []string {
	ret := (*st)[len(*st)-n:]
	(*st) = (*st)[:len(*st)-n]
	return ret
}

func (st *stack) drop(d []string) {
	(*st) = append((*st), d...)
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	stacks := []stack{}
	for i := 0; i < NUM_STACKS; i++ {
		stacks = append(stacks, stack{})
	}

	var instructions []instruction

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "[") {
			//    [G] [R]             [P]
			//    [H] [W]     [T] [P] [H]
			//    [F] [T] [P] [B] [D] [N]
			//[L] [T] [M] [Q] [L] [C] [Z]
			//[C] [C] [N] [V] [S] [H] [V] [G]
			for i := 0; i < len(scanner.Text()); i++ {
				index := (i * 4) + 1
				if index < len(scanner.Text()) {
					ch := string(scanner.Text()[index])
					if ch != " " {
						stacks[i].headPush(ch)
					}
				}
			}
		}

		if strings.Contains(scanner.Text(), "move") {
			// e.g. move 5 from 8 to 2
			parts := strings.Split(scanner.Text(), " ")
			quantity, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])
			instructions = append(instructions, instruction{quantity: quantity, from: from, to: to})
		}
	}
	part1(copyStack(stacks), instructions)
	part2(copyStack(stacks), instructions)
}

func part2(stacks []stack, instructions []instruction) {
	for _, ins := range instructions {
		// Correct for the off by one numbering
		to := ins.to - 1
		from := ins.from - 1
		// fmt.Printf("\t Start From: %s\n", stacks[from])
		// fmt.Printf("\t Start To: %s\n", stacks[to])
		// fmt.Printf("\t Instruction: %s\n", ins)

		crates := stacks[from].pick(ins.quantity)
		stacks[to].drop(crates)

		// fmt.Printf("\t End From: %s\n", stacks[from])
		// fmt.Printf("\t End To: %s\n", stacks[to])
		// fmt.Printf("\n")
	}

	out := ""
	for _, s := range stacks {
		out += s.pop()
	}
	fmt.Println(out)
}

func part1(stacks []stack, instructions []instruction) {
	for _, ins := range instructions {
		// Correct for the off by one numbering
		to := ins.to - 1
		from := ins.from - 1
		// fmt.Printf("\t Start From: %s\n", stacks[from])
		// fmt.Printf("\t Start To: %s\n", stacks[to])
		// fmt.Printf("\t Instruction: %s\n", ins)

		for i := 0; i < ins.quantity; i++ {
			crate := stacks[from].pop()
			stacks[to].push(crate)
		}

		// fmt.Printf("\t End From: %s\n", stacks[from])
		// fmt.Printf("\t End To: %s\n", stacks[to])
		// fmt.Printf("\n")
	}

	out := ""
	for _, s := range stacks {
		out += s.pop()
	}
	fmt.Println(out)
}

func copyStack(stacks []stack) []stack {
	var out []stack
	for _, s := range stacks {
		var new stack
		new = append(new, s...)
		out = append(out, new)
	}
	return out
}
