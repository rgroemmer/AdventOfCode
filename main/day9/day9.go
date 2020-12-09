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
	//numberList := []int{35,20,15,25,47,40,62,55,65,95,102,117,150,182,127,219,299,277,309,576}

	var j int
	var magicNumb int
	for i := 5; i < len(numberList)-1; i++ {
		preamble := numberList[j:i]
		checkValue := numberList[i]
		if !checkNumbs(preamble, checkValue) {
			magicNumb = checkValue
		}
		j++
	}

	y := 0
	for i := 5; i < len(numberList)-1; i++ {
		preamble := numberList[y:i]
		subset := findSubset(preamble, magicNumb)
		if subset != nil && len(subset) > 1 {
			sort.Ints(subset)
			println("part2: ", subset[0]+subset[len(subset)-1])
		}

		y++
	}
	println("magic Number: ", magicNumb)
}
