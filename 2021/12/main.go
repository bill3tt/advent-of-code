package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type location struct {
	name  string
	paths []*location
}

func (l location) small() bool {
	for _, r := range l.name {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (l location) String() string {
	var destinations []string
	for _, d := range l.paths {
		destinations = append(destinations, d.name)
	}
	return fmt.Sprintf("%s -> %s", l.name, strings.Join(destinations, " "))
}

var paths [][]string

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	locations := map[string]*location{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		fromName, toName := parts[0], parts[1]

		from, ok := locations[fromName]
		if !ok {
			from = &location{name: fromName, paths: []*location{}}
			locations[fromName] = from
		}

		to, ok := locations[toName]
		if !ok {
			to = &location{name: toName, paths: []*location{}}
			locations[toName] = to
		}
		from.paths = append(from.paths, to)
		to.paths = append(to.paths, from)
	}

	var current []string
	visited := map[*location]int{}

	start := locations["start"]

	// Generate all paths
	dfs(start, visited, current)

	for _, p := range paths {
		fmt.Println(p)
	}
	fmt.Printf("Total paths: %d\n", len(paths))
}

func dfs(loc *location, visited map[*location]int, current []string) {
	fmt.Printf("Current path: %s, Current location: %s\n", current, loc.name)
	if loc.name == "end" {
		current = append(current, loc.name)
		fmt.Printf("Found path: %s\n", current)
		paths = append(paths, current)
		return
	}

	visited[loc]++

	if !proceed(visited) {
		// Progress no further
		fmt.Printf("Progressing no further. %d\n", visited[loc])
		return
	}

	for _, d := range loc.paths {
		vis := copyMap(visited)
		p := copySlice(current)
		p = append(p, loc.name)
		dfs(d, vis, p)
	}
}

func proceed(visited map[*location]int) bool {
	// Is this proposed path valid to proceed along?
	//   * You can only visit one small cave twice
	twiceVisitedCount := 0

	for l, v := range visited {
		if (l.name == "start" || l.name == "end") && v > 1 {
			return false
		}
		if l.small() && v > 2 {
			return false
		}
		if l.small() && v > 1 {
			twiceVisitedCount++
		}
	}

	return twiceVisitedCount < 2
}

func copyMap(in map[*location]int) map[*location]int {
	out := make(map[*location]int)
	for k, v := range in {
		out[k] = v
	}
	return out
}

func copySlice(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	return out
}
