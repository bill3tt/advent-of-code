package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	// 123
	// 234
	//
	// 1234
	// 1231

	total := 0
	cals := []int{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			cals = append(cals, total)
			total = 0
		}
		val, _ := strconv.Atoi(scanner.Text())
		total += val
	}
	cals = append(cals, total)

	fmt.Println(cals)
	// Find max of calories

	sort.Ints(cals)

	fmt.Printf("Max calories is: %d\n", cals[len(cals)-1])

	fmt.Printf("Max calories of top 3 is: %d\n", cals[len(cals)-1]+cals[len(cals)-2]+cals[len(cals)-3])
}
