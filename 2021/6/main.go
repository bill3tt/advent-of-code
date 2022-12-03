package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var fish []int
	for _, s := range strings.Split(scanner.Text(), ",") {
		f, _ := strconv.Atoi(s)
		fish = append(fish, f)
	}

	naiveFish := make([]int, len(fish))
	copy(naiveFish, fish)
	fmt.Println(naive(naiveFish, 80))

	optimizedFish := make([]int, len(fish))
	copy(optimizedFish, fish)
	fmt.Println(optimized(optimizedFish, 80))
	fmt.Println(optimized(optimizedFish, 256))
}

func naive(fish []int, days int) int {
	day := 1
	for day <= days {
		var newFish []int

		for i := range fish {
			// Firstly decrement the timer
			fish[i]--

			// Don't add the new fish to the array right away
			// to avoid getting processed in this iteration
			if fish[i] == -1 {
				fish[i] = 6
				newFish = append(newFish, 8)
			}
		}
		fish = append(fish, newFish...)
		day++
	}
	return len(fish)
}

func optimized(fish []int, days int) int {
	// Firstly convert slice of ints into count of fish at a given day
	count := make([]int, 9)
	// [0, 1, 1, 2, 0, 2] -> [2, 2, 2]
	for _, f := range fish {
		count[f]++
	}

	day := 1

	for day <= days {
		expiring := count[0]

		// Truncate first element of slice
		count = count[1:]
		// Schedule new fish to be created
		count = append(count, expiring)
		// Schedule rebirth of fish
		count[6] += expiring

		day++
	}

	return sum(count)
}

func sum(count []int) int {
	total := 0
	for _, c := range count {
		total += c
	}
	return total
}
