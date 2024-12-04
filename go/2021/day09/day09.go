package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strconv"
)

type coodinate struct {
	row int
	col int
}

func part1(input []string) (result int) {
	heatmap := parseInput(input)
	var results []int

	for row := 0; row < len(heatmap); row++ {
		for col := 0; col < len(heatmap[row]); col++ {
			//get adjacent
			adjacent := getAdjacentFieldIndexes(heatmap, row, col)
			depth := heatmap[row][col]

			var count int
			for _, c := range adjacent {
				if depth < heatmap[c.row][c.col] {
					count++
				}
			}
			if count == len(adjacent) {
				results = append(results, depth)
			}
		}
	}

	return
}

func getAdjacentFieldIndexes(heatmap [][]int, row int, col int) (adjacent []coodinate) {
	if row == 0 {
		if col == 0 {
			adjacent = append(adjacent, coodinate{
				row: 0,
				col: 1,
			}, coodinate{
				row: 1,
				col: 0,
			})
			return adjacent

		}
		if col == len(heatmap[row])-1 {
			adjacent = append(adjacent, coodinate{
				row: row + 1,
				col: col,
			}, coodinate{
				row: row,
				col: col - 1,
			})
			return adjacent
		}
		adjacent = append(adjacent, coodinate{
			row: row,
			col: col - 1,
		}, coodinate{
			row: row,
			col: col + 1,
		}, coodinate{
			row: row + 1,
			col: col,
		})
	}

	if row == len(heatmap)-1 {
		if col == 0 {
			adjacent = append(adjacent, coodinate{
				row: row - 1,
				col: col,
			}, coodinate{
				row: row,
				col: col + 1,
			})
			return adjacent
		}

		if col == len(heatmap[row])-1 {
			adjacent = append(adjacent, coodinate{
				row: row - 1,
				col: col,
			}, coodinate{
				row: row,
				col: col - 1,
			})
			return adjacent
		}

		adjacent = append(adjacent, coodinate{
			row: row,
			col: col - 1,
		}, coodinate{
			row: row,
			col: col + 1,
		}, coodinate{
			row: row - 1,
			col: col,
		})
	}

	adjacent = append(adjacent, coodinate{
		row: row + 1,
		col: col,
	}, coodinate{
		row: row - 1,
		col: col,
	}, coodinate{
		row: row,
		col: col + 1,
	}, coodinate{
		row: row,
		col: col - 1,
	})

	return adjacent
}

func parseInput(input []string) [][]int {
	heatmap := make([][]int, len(input))
	for i, line := range input {
		heatmap[i] = make([]int, len(line))
		for j, char := range line {
			heatmap[i][j], _ = strconv.Atoi(string(char))
		}
	}
	return heatmap
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
