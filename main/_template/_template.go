package main

import (
	helper "../helper"
	"fmt"
	"os"
)

func part1(input string)  (result int){

	return
}

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day9/input.txt")
	if err != nil {
		println("Error: ", err)
	}
	input := helper.ReadeFile(r)
	fmt.Println(input)
}
