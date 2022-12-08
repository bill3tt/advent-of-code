package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const KIND_DIR = "directory"
const KIND_FILE = "file"

const LIST = "ls"
const CHDIR = "cd"

var root *Node

type Node struct {
	path, kind, name string
	children         []*Node
	parent           *Node
	size             int
}

func (n *Node) getSize() int {
	if n.kind == KIND_DIR && n.size == 0 {
		total := 0
		for _, c := range n.children {
			total = total + c.getSize()
		}
		n.size = total
	}
	return n.size
}

type Command struct {
	command  string
	argument string
	output   []string
}

func (c Command) String() string {
	return fmt.Sprintf("c: %s a: %s", c.command, c.argument)
}

type nodeStack []*Node

func (ns *nodeStack) push(n *Node) {
	(*ns) = append((*ns), n)
}

func (ns *nodeStack) pop() *Node {
	o := (*ns)[0]
	(*ns) = (*ns)[1:]
	return o
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	commands := parse(input)

	root = &Node{path: "", name: "", kind: KIND_DIR, children: []*Node{}, parent: nil}
	pwd := root

	// Process the full set of commands to build the directory tree
	for _, c := range commands {
		switch c.command {
		case CHDIR:
			pwd = chdir(pwd, c)
		case LIST:
			pwd = ls(pwd, c)
		}
	}

	// Part 1
	total := 0
	dirs := dirs(root)
	for _, d := range dirs {
		if d.getSize() < 100000 {
			total += d.getSize()
		}
	}
	fmt.Printf("Part 1: %d\n", total)

	// Part 2
	total_space := 70000000
	required := 30000000
	free_space := total_space - root.getSize()
	target := required - free_space
	fmt.Printf("Total %d, Free: %d, Required: %d, Target: %d\n", total_space, free_space, required, target)

	min := free_space
	for _, d := range dirs {
		if d.getSize() > target && d.getSize() < min {
			min = d.getSize()
		}
	}
	fmt.Printf("Part 2: %d\n", min)
}

// dirs returns a slice of all directories in the file system
func dirs(root *Node) []*Node {
	stack := nodeStack{root}
	dirs := []*Node{}
	for len(stack) > 0 {
		n := stack.pop()
		for _, c := range n.children {
			if c.kind == KIND_DIR {
				stack.push(c)
				dirs = append(dirs, c)
			}
		}
	}
	return dirs
}

func chdir(pwd *Node, command Command) *Node {
	children := map[string]*Node{}
	for _, child := range pwd.children {
		children[child.name] = child
	}
	switch command.argument {
	case "..":
		return pwd.parent
	case "/":
		return root
	default:
		if _, ok := children[command.argument]; !ok {
			panic(fmt.Sprintf("%s does not exist in %s", command.argument, pwd.path))
		}
		return children[command.argument]
	}
}

func ls(pwd *Node, command Command) *Node {
	var nodes []*Node
	for _, o := range command.output {
		node := build(o, pwd)
		nodes = append(nodes, node)
	}
	pwd.children = append(pwd.children, nodes...)
	return pwd
}

// build builds a node entry based on the output of an ls command
func build(line string, pwd *Node) *Node {
	parts := strings.Split(line, " ")
	if strings.Contains(line, "dir") {
		return &Node{
			path:     pwd.path + "/" + parts[1],
			name:     parts[1],
			parent:   pwd,
			kind:     KIND_DIR,
			children: []*Node{},
		}
	}
	size, _ := strconv.Atoi(parts[0])
	return &Node{
		path:     pwd.path + "/" + parts[1],
		name:     parts[1],
		parent:   pwd,
		kind:     KIND_FILE,
		children: nil,
		size:     size,
	}
}

func parse(input []string) []Command {
	var commands []Command
	i := 0
	for i < len(input) {
		// First item is a command rather than output.
		parts := strings.Split(input[i], " ")
		var command, argument string
		command = parts[1]
		if len(parts) > 2 {
			argument = parts[2]
		}
		var output []string

		var j int
		for j = i + 1; j < len(input) && input[j][0] != byte('$'); j++ {
			output = append(output, input[j])
		}
		i = j
		commands = append(commands, Command{command: command, argument: argument, output: output})
	}
	return commands
}
