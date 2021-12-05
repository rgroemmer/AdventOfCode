package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	startx, starty, endx, endy int
}

func parseInput(input []string) (lines []line) {
	for _, s := range input {
		split := strings.Split(s, ",")

		startx, _ := strconv.Atoi(split[0])
		starty, _ := strconv.Atoi(split[1])
		endx, _ := strconv.Atoi(split[2])
		endy, _ := strconv.Atoi(split[3])

		lines = append(lines, line{
			startx: startx,
			starty: starty,
			endx:   endx,
			endy:   endy,
		})
	}
	return
}

func part1(input []string) (result int) {
	lines := parseInput(input)
	var field [10][10]int

	for _, line := range lines {
		if line.startx == line.endx || line.starty == line.endy {
			drawLine(line, &field)
			printField(field)
		}
	}

	return
}

func printField(field [10][10]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			value := field[j][i]
			fmt.Print(value)
		}
		println("")
	}
	fmt.Println("#-----#-----#")
}

func drawLine(l line, field *[10][10]int) {
	if l.startx == l.endx {
		lowValue := l.starty
		highValue := l.endy

		if l.starty > l.endy {
			highValue = l.starty
			lowValue = l.endy
		}
		//	vertical
		for y := lowValue; y <= highValue; {
			field[l.startx][y]++
			y++
		}
	}
	if l.starty == l.endy {
		lowValue := l.startx
		highValue := l.endx

		if l.startx > l.endx {
			highValue = l.startx
			lowValue = l.endx
		}

		//	horiontal
		for x := lowValue; x <= highValue; {
			field[x][l.starty]++
			x++
		}
	}
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
	//fmt.Println(part2(input))
}
