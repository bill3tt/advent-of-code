package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// 199  A        0, 3 -> the rest are in both sets
// 200  A B      1, 4
// 208  A B C
// 210    B C D
// 200  E   C D
// 207  E F   D
// 240  E F G
// 269    F G H
// 260      G H
// 263        H

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var input []int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val, err := strconv.Atoi((scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, val)
	}

	count := 0

	for i := 3; i < len(input); i++ {
		if math.Abs(float64(input[i])) > math.Abs(float64(input[i-3])) {
			count++
		}
	}
	fmt.Println(count)
}
