package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x, y int
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

func part1(input []string) (result int) {
	grid := parseInput(input)
	for step := 0; step < 100; step++ {
		flashed := make(map[point]bool)
		for x := range grid {
			for y := range grid[x] {
				grid[x][y]++
				if grid[x][y] > 9 {
					flashOctopus(grid, point{x, y}, flashed)
				}
			}
		}
		result += len(flashed)
		for k, _ := range flashed {
			grid[k.x][k.y] = 0
		}
	}
	return
}

func part2(input []string) (result int) {
	grid := parseInput(input)
	for step := 0; true; step++ {
		flashed := make(map[point]bool)
		for x := range grid {
			for y := range grid[x] {
				grid[x][y]++
				if grid[x][y] > 9 {
					flashOctopus(grid, point{x, y}, flashed)
				}
			}
		}
		for k, _ := range flashed {
			grid[k.x][k.y] = 0
		}
		if checkInSync(grid) {
			result = step + 1
			break
		}
	}
	return
}

func checkInSync(grid [][]int) bool {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != 0 {
				return false
			}
		}
	}
	return true
}

func flashOctopus(grid [][]int, p point, flashed map[point]bool) {
	//check if octopus has already flashed
	for k, _ := range flashed {
		if k == p {
			return
		}
	}
	//set flashing octopus to 0
	flashed[p] = true

	adjacent := getAdjacent(p)
	//increment all adjacent by one
	for _, ap := range adjacent {
		grid[ap.x][ap.y]++
		if grid[ap.x][ap.y] > 9 {
			flashOctopus(grid, ap, flashed)
		}
	}
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

func getAdjacent(p point) (adjacent []point) {
	for i := p.x - 1; i < p.x+2; i++ {
		for j := p.y - 1; j < p.y+2; j++ {
			if i == p.x && j == p.y {
				continue
			}
			if (i >= 0 && j >= 0) && (i < 10 && j < 10) {
				adjacent = append(adjacent, point{x: i, y: j})
			}
		}
	}
	return
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
