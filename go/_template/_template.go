package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
)

func part1(input []string)  (result int) {

	return
}

func part2(input []string)  (result int) {

	return
}

func main() {
	r, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input, err := helper.InputToStringSlice(string(r))
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
