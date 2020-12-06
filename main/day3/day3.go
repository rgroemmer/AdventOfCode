package main

import (
	helper "../helper"
	"os"
)

func movePosition(length, move, actPosition int) int {
	actPosition += move
	if actPosition > length-1 {
		actPosition -= length
	}
	return actPosition
}

func getTreeCount(treeMap []string, moveRight int, jumpTwice bool) int {
	var pos int
	var count int
	for i, line := range treeMap {
		if !jumpTwice {
			pos = movePosition(len(line), moveRight, pos)
			if line[pos] == '#' {
				count++
			}
		} else {
			if i%2 == 0 {
				pos = movePosition(len(line), moveRight, pos)
				if line[pos] == '#' {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day3/input.txt")
	if err != nil {
		panic(err)
	}

	treeMap := helper.ReadeFileByLine(r)
	treeMap = treeMap[1:]

	r1d1 := getTreeCount(treeMap, 1, false)
	r3d1 := getTreeCount(treeMap, 3, false)
	r5d1 := getTreeCount(treeMap, 5, false)
	r7d1 := getTreeCount(treeMap, 7, false)
	treeMap = treeMap[1:]
	r1d2 := getTreeCount(treeMap, 1, true)

	println("r1d1: ", r1d1)
	println("r3d1: ", r3d1)
	println("r5d1: ", r5d1)
	println("r7d1: ", r7d1)
	println("r1d2: ", r1d2)
	println("magic number: ", r1d1*r3d1*r5d1*r7d1*r1d2)
}
