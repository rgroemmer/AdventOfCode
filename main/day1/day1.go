package main

import (
	helper "../helper"
	"os"
)

// BenchmarkDynamicProgramming approach for efficiency
func IsSubsetSum(set []int, n, sum int) bool {
	subset := make([][]bool, n+1)
	for index := range subset {
		subset[index] = make([]bool, sum+1)
	}
	for i := 0; i <= n; i++ {
		subset[i][0] = true
	}
	for i := 0; i <= sum; i++ {
		subset[0][i] = false
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j < set[i-1] {
				subset[i][j] = subset[i-1][j]
			}
			if j >= set[i-1] {
				subset[i][j] = subset[i-1][j] || subset[i-1][j-set[i-1]]
			}
		}
	}
	return subset[n][sum]
}

func Numb2(expenses []int) int {
	var result int
	for _, value1 := range expenses {
		for _, value2 := range expenses {
			if (value1 + value2) == 2020 {
				result = value1 * value2
			}
		}
	}
	return result
}

func Numb3(expenses []int) int {
	var result int
	for _, value1 := range expenses {
		for _, value2 := range expenses {
			for _, value3 := range expenses {
				if (value1 + value2 + value3) == 2020 {
					result = value1 * value2 * value3
				}
			}
		}
	}
	return result
}

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day1/input.txt")
	if err != nil {
		panic(err)
	}

	expenses, err := helper.ReadInts(r)
	if err != nil {
		panic(err)
	}

	println("Product of two Numbers that equal: ", Numb2(expenses))
	println("Product of three Numbers that equal: ", Numb3(expenses))
}
