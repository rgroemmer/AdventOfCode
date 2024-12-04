package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
)

func main() {
	r, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input, err := helper.InputToIntSlice(string(r), ",")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []int) (result int) {
	count := 0
	for s := 1; s < len(input); s++ {
		if isIncreasing(input[s-1], input[s]) {
			count++
		}
	}
	return count
}

func part2(input []int) (result int) {
	var newInput []int

	for s := 2; s < len(input); s++ {
		window := input[s-2] + input[s-1] + input[s]
		newInput = append(newInput, window)
	}
	return part1(newInput)
}

func isIncreasing(prev, now int) bool {
	if prev < now {
		return true
	}
	return false
}
