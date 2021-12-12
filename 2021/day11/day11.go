package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strconv"
)

type coordinate struct {
	row int
	col int
}

func parseInput(input []string) [][]int {
	grid := make([][]int, len(input))
	for i, line := range input {
		grid[i] = make([]int, len(input[0]))
		for j, char := range line {
			grid[i][j], _ = strconv.Atoi(string(char))
		}
	}
	return grid
}

func getAdjacent(row, col int) (adjacent []coordinate) {
	for i := row - 1; i < row+2; i++ {
		for j := col - 1; j < col+2; j++ {
			if i == row && j == col {
				continue
			}
			if (i >= 0 && j >= 0) && (i < 5 && j < 5) {
				adjacent = append(adjacent, coordinate{row: i, col: j})
			}
		}
	}
	return
}

func incrementAll(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			slice[i][j]++
		}
	}
}
func getFlashingOctos(grid [][]int) []coordinate {
	var flashingList []coordinate
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > 9 {
				flashingList = append(flashingList, coordinate{row: i, col: j})
			}
		}
	}
	return flashingList
}

func part1(input []string) (result int) {
	grid := parseInput(input)

	for step := 0; step < 10; step++ {

		//increase all by one
		incrementAll(grid)

		//get octos who will flash
		var flashingList []coordinate
		flashingList = getFlashingOctos(grid)

		for len(flashingList) != 0 {
			//do flash for actual octo
			doFlash(grid, flashingList[0])
			result++
			//retrieve new flashing list and overwrite old
			flashingList = getFlashingOctos(grid)
		}

		printGrid(grid)
	}

	return
}

func doFlash(grid [][]int, c coordinate) {
	//set flashing octopus to 0
	grid[c.row][c.col] = 0
	//var nextFlashes []coordinate

	adjacent := getAdjacent(c.row, c.col)
	//increment all adjacent by one
	for _, a := range adjacent {
		if grid[a.row][a.col] != 0 {
			grid[a.row][a.col]++
		}
	}
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func part2(input []string) (result int) {

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
