package main

import (
	helper "../helper"
	"os"
	"sort"
)

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day9/input.txt")
	if err != nil {
		println("Error: ", err)
	}
	numberList, _ := helper.ReadeFile(r)

}
