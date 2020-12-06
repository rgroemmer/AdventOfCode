package main

import (
	helper "../helper"
	"os"
	"sort"
)

func getInt(binaryString string) int {
	var result int
	for _, char := range binaryString {
		if char == 'F' || char == 'L' {
			result <<= 1
		}
		if char == 'B' || char == 'R' {
			result <<= 1
			result += 1
		}
	}
	return result
}

func getSeatId(row, col int) int {
	return row*8 + col
}

func getYourSeat(ids []int) int {
	value := ids[0]
	for i := 0; i < len(ids); i++ {
		if !(ids[i] == value) {
			if ids[i-1] == value-1 && ids[i] == value+1 {
				return value
			}
			value++
		}
		value++
	}
	return -1
}

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day5/input.txt")
	if err != nil {
		println("Error: ", err)
	}
	seatList := helper.ReadeFileByLine(r)
	var ids []int
	for _, seat := range seatList {
		id := getSeatId(getInt(seat[:7]), getInt(seat[7:]))
		ids = append(ids, id)
	}
	sort.Ints(ids)
	println("Magic number: ", ids[len(ids)-1])
	println("Your seat: ", getYourSeat(ids))
}
