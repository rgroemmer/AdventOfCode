package main

import (
	helper "../helper"
	"os"
	"sort"
)

func checkNumbs(preamble []int, checkValue int) bool {
	for _, numb1 := range preamble {
		for _, numb2 := range preamble {
			if (numb1+numb2 == checkValue) && (numb1 != numb2) {
				return true
			}
		}
	}
	return false
}

func findSubset(preamble []int, checkValue int) []int {
	aktNumb := 0
	var array []int
	for i := 0; i < len(preamble); i++ {
		aktNumb += preamble[i]
		array = append(array, preamble[i])
		if aktNumb > checkValue {
			array = nil
			aktNumb = 0
		}
		if aktNumb == checkValue {
			return array
		}
	}
	return nil
}

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day9/input.txt")
	if err != nil {
		println("Error: ", err)
	}
	numberList, _ := helper.ReadInts(r)

	var j int
	part1 := 0
	for i := 25; i < len(numberList)-1; i++ {
		preamble := numberList[j:i]
		checkValue := numberList[i]
		if !checkNumbs(preamble, checkValue) {
			part1 = checkValue
		}
		j++
	}
	println("part1: ", part1)

	y := 0
	for i := 25; i < len(numberList)-1; i++ {
		preamble := numberList[y:i]
		subset := findSubset(preamble, part1)
		if subset != nil && len(subset) > 1 {
			sort.Ints(subset)
			println("part2: ", subset[0]+subset[len(subset)-1])
		}
		y++
	}
}
