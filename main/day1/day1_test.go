package main

import (
	helper "../helper"
	"os"
	"testing"
)

func BenchmarkDynamicProgramming(b *testing.B) {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/day1/input.txt")
	if err != nil {
		panic(err)
	}

	expenses, err := helper.ReadInts(r)
	if err != nil {
		panic(err)
	}

	for n := 0; n < 100; n++ {
		IsSubsetSum(expenses, len(expenses), 2020)
	}
}

func BenchmarkForLoop(b *testing.B) {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/day1/input.txt")
	if err != nil {
		panic(err)
	}

	expenses, err := helper.ReadInts(r)
	if err != nil {
		panic(err)
	}

	for n := 0; n < 100; n++ {
		Numb2(expenses)
	}
}
