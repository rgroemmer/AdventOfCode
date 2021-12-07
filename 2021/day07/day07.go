package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
)

func part1(input []int) (result int) {
	_, max := MinMax(input)
	usedFuel := make([]int, max+1)

	for i := 0; i <= max; i++ {
		for _, crab := range input {
			cost := calcFuelToPos(crab, i)
			if cost > 0 {
				usedFuel[i] += cost
			}
		}
	}

	min, _ := MinMax(usedFuel)

	return min
}

func part2(input []int) (result int) {
	_, max := MinMax(input)
	usedFuel := make([]int, max+1)

	for i := 0; i <= max; i++ {
		for _, crab := range input {
			cost := calcFuelToPos(crab, i)
			if cost > 0 {
				usedFuel[i] += calcGaussSum(cost)
			}
		}
	}

	min, _ := MinMax(usedFuel)

	return min
}

func calcGaussSum(cost int) int {
	return ((cost * cost) + cost) / 2
}

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

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if value == 0 {
			continue
		}
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func calcFuelToPos(crab, pos int) int {
	var cost int
	if crab < pos {
		cost = pos - crab
	} else {
		cost = crab - pos
	}
	return cost
}
