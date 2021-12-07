package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
)

func main() {
	r, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input, err := helper.InputToIntSlice(string(r), ",")
	if err != nil {
		os.Exit(1)
	}
	//fmt.Println(part1(input, 80))
	fmt.Println(part2(input, 256))
}

func part1(input []int, days int) (result int) {
	actFishes := input
	var newFishes []int

	for day := 1; day <= days; day++ {
		for j := 0; j < len(actFishes); j++ {
			actFish := &actFishes[j]

			if *actFish == 0 {
				//new fish is born
				newFishes = append(newFishes, 8)
				*actFish = 6
				continue
			}

			if *actFish > 0 {
				*actFish--
			}
		}
		actFishes = append(actFishes, newFishes...)
		newFishes = nil
	}
	return len(actFishes)
}

func countFishs(fishlevels []int) (count int) {
	for _, level := range fishlevels {
		count += level
	}
	return
}

func part2(input []int, days int) (result int) {
	fishLevels := make([]int, 9)
	newFishes := make([]int, 9)
	for _, fish := range input {
		fishLevels[fish]++
	}

	for day := 1; day <= days; day++ {
		for i := 0; i < 9; i++ {
			if fishLevels[i] > 0 {
				if i == 0 {
					newFishes[6] = fishLevels[0]
					newFishes[8] = fishLevels[0]
				} else {
					newFishes[i-1] += fishLevels[i]
				}

			}
		}
		fishLevels = newFishes
		newFishes = make([]int, 9)
	}
	return countFishs(fishLevels)
}
