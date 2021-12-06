package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ventLine struct {
	start point
	end   point
}

type point struct {
	x int
	y int
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

	fmt.Println(part1(input, 1000))
	fmt.Println(part2(input, 1000))
}

func part1(input []string, fieldsize int) (result int) {
	lines := parseInput(input)
	field := make([][]int, fieldsize)
	for i := range field {
		field[i] = make([]int, fieldsize)
	}

	for _, line := range lines {
		if line.start.x == line.end.x || line.start.y == line.end.y {
			drawLine(line, field)
		}
	}

	return countPointsHigherThan(field, 2)
}

func part2(input []string, fieldsize int) (result int) {
	lines := parseInput(input)
	field := make([][]int, fieldsize)
	for i := range field {
		field[i] = make([]int, fieldsize)
	}

	for _, line := range lines {
		drawLine(line, field)
	}

	return countPointsHigherThan(field, 2)
}

func parseInput(input []string) (lines []ventLine) {
	for _, s := range input {
		split := strings.Split(s, ",")
		points := make([]int, 4)

		for i, value := range split {
			points[i], _ = strconv.Atoi(value)
		}

		lines = append(lines, ventLine{
			start: point{
				x: points[0],
				y: points[1],
			},
			end: point{
				x: points[2],
				y: points[3],
			},
		})
	}
	return
}

func countPointsHigherThan(field [][]int, limit int) int {
	var count int
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field); j++ {
			if field[i][j] >= limit {
				count++
			}
		}
	}

	return count
}

//only for tests
func printField(field [][]int) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field); j++ {
			value := field[j][i]
			fmt.Print(value)
		}
		println("")
	}
	fmt.Println("#-----#-----#")
}

func getPointsBetween(l ventLine) [][]int {
	var betweenX []int
	var betweenY []int

	for x := l.start.x; ; {
		betweenX = append(betweenX, x)
		if x == l.end.x {
			break
		}

		if l.start.x > l.end.x {
			x--
		} else {
			x++
		}

	}

	for y := l.start.y; ; {
		betweenY = append(betweenY, y)
		if y == l.end.y {
			break
		}
		if l.start.y > l.end.y {
			y--
		} else {
			y++
		}

	}

	return [][]int{betweenX, betweenY}
}

func drawLine(l ventLine, field [][]int) {
	if l.start.x == l.end.x {
		lowValue := l.start.y
		highValue := l.end.y

		if l.start.y > l.end.y {
			highValue = l.start.y
			lowValue = l.end.y
		}
		//	vertical
		for y := lowValue; y <= highValue; y++ {
			field[l.start.x][y]++
		}
	} else if l.start.y == l.end.y {
		lowValue := l.start.x
		highValue := l.end.x

		if l.start.x > l.end.x {
			highValue = l.start.x
			lowValue = l.end.x
		}

		//	horiontal
		for x := lowValue; x <= highValue; x++ {
			field[x][l.start.y]++
		}
	} else {
		//TODO getNumbersBetwwen
		points := getPointsBetween(l)
		for i := 0; i < len(points[0]); i++ {
			x := points[0][i]
			y := points[1][i]

			field[x][y]++
		}
	}
}
