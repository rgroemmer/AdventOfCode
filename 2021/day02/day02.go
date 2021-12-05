package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	depth      int
	horizontal int
	aim        int
}

func part1(input []string) (result int) {
	actPosiston := position{}
	for _, line := range input {
		split := strings.Split(line, " ")
		command := split[0]
		steps, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}
		switch command {
		case "forward":
			actPosiston.horizontal += steps
		case "up":
			actPosiston.depth -= steps
		case "down":
			actPosiston.depth += steps

		}
	}
	return actPosiston.horizontal * actPosiston.depth
}

func part2(input []string) (result int) {
	actPosiston := position{
		depth:      0,
		horizontal: 0,
		aim:        0,
	}
	for _, line := range input {
		split := strings.Split(line, " ")
		command := split[0]
		steps, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}
		switch command {
		case "forward":
			actPosiston.horizontal += actPosiston.horizontal + steps
			actPosiston.depth += actPosiston.aim * steps
		case "up":
			actPosiston.aim = actPosiston.aim - steps
		case "down":
			actPosiston.aim = actPosiston.aim + steps
		}
	}
	return actPosiston.horizontal * actPosiston.depth
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
