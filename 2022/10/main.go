package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name  string
	value int
}

type Running struct {
	Instruction
	start int
}

type CPU struct {
	instructions []Instruction
	icounter     int
	current      Running
	next         int
	register     int
	cycle        int
	pixels       string
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (c *CPU) tick() {
	c.cycle++

	// If we don't have a running command, load one.
	if !c.running() {
		c.current = Running{Instruction: c.instructions[c.icounter], start: c.cycle}
		c.icounter++
	}

	// Do we have the output of a previous command to add to the register?
	if c.next != 0 {
		c.register = c.register + c.next
		c.next = 0
	}

	switch c.current.name {
	case "noop":
		c.noop()
	case "addx":
		c.addx()
	}

	c.draw()
}

func (c *CPU) draw() {
	// Do we need to draw a newline?
	if (c.cycle-1)/40 != (c.cycle-2)/40 {
		c.pixels = c.pixels + "\n"
	}

	// pixel = 0 .. 39
	pixel := (c.cycle - 1) % 40
	sprite := c.register

	// Is the pixel we are drawing and the sprite within one character of each other?
	draw := abs(pixel-sprite) < 2

	char := "."
	if draw {
		char = "#"
	}
	c.pixels = c.pixels + char
}

func (c *CPU) addx() {
	if c.cycle > c.current.start {
		c.next = c.current.value
		c.unload()
	}
}

func (c *CPU) noop() {
	c.unload()
}

func (c *CPU) unload() {
	c.current = Running{}
}

func (c *CPU) running() bool {
	return c.current != Running{}
}

func (c CPU) signalValue() int {
	return c.cycle * c.register
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var instructions []Instruction
	for scanner.Scan() {
		parts := make([]string, 2)
		copy(parts, strings.Split(scanner.Text(), " "))
		v, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, Instruction{parts[0], v})
	}

	cpu := CPU{instructions: instructions, register: 1}
	total := 0
	for cpu.icounter < len(cpu.instructions) {
		cpu.tick()
		fmt.Printf("Cycle: %d, Reg: %d, Signal: %d, Pixel: %d, Line: %d\n", cpu.cycle, cpu.register, cpu.signalValue(), (cpu.cycle-1)%40, (cpu.cycle-1)/40)
		if (cpu.cycle-20)%40 == 0 {
			total += cpu.signalValue()
		}
	}
	fmt.Println(total)
	fmt.Println(cpu.pixels)
}
