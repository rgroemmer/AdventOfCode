package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strconv"
)

func isMostCommonBitHigh(input []string,pos int) bool {
	var highcount int
	var lowcount int

	for _, number := range input {
		runes := []rune(number)
		actNumb := runes[pos]

		if actNumb == '1'{
			highcount++
		}
		if actNumb == '0'{
			lowcount++
		}
	}
	return highcount >= lowcount
}

func isLeastCommonBitHigh(input []string,pos int) bool {
	var highcount int
	var lowcount int

	for _, number := range input {
		runes := []rune(number)
		actNumb := runes[pos]

		if actNumb == '1'{
			highcount++
		}
		if actNumb == '0'{
			lowcount++
		}
	}
	return highcount < lowcount
}

func isBitAtPosHigh(numb string, pos int)  bool{
	runes := []rune(numb)
	actNumb := runes[pos]
	return actNumb == '1'
}

func part1(input []string)  (result int){
	var gamma int
	var epsilon int

	for i := 0; i < len(input[0]); i++ {
		gamma <<= 1
		epsilon <<= 1
		if isMostCommonBitHigh(input, i) {
			gamma += 1
		} else {
			epsilon += 1
		}
	}
	return gamma * epsilon
}

func part2(input []string)  (result int64){
	var temp []string
	var oxy int64
	var c02 int64

	filterd := input
	for i := 0; i < len(input[0]); i++ {
		commonIsHigh := isMostCommonBitHigh(filterd, i)
		temp = nil
		for j := 0; j < len(filterd); j++ {
			actNumb := filterd[j]
			if commonIsHigh {
				if isBitAtPosHigh(actNumb, i) {
					temp = append(temp, actNumb)
				}
			}
			if !commonIsHigh {
				if !isBitAtPosHigh(actNumb, i) {
					temp = append(temp, actNumb)
				}
			}
		}
		filterd = temp
		if len(filterd) == 1 {
			oxy, _ = strconv.ParseInt(filterd[0], 2, 64)
		}
	}

	filterd = input
	for i := 0; i < len(input[0]); i++ {
		leastIsHigh := isLeastCommonBitHigh(filterd, i)
		temp = nil
		for j := 0; j < len(filterd); j++ {
			actNumb := filterd[j]
			if leastIsHigh {
				if isBitAtPosHigh(actNumb, i) {
					temp = append(temp, actNumb)
				}
			}
			if !leastIsHigh {
				if !isBitAtPosHigh(actNumb, i) {
					temp = append(temp, actNumb)
				}
			}
		}
		filterd = temp
		if len(filterd) == 1 {
			c02, _ = strconv.ParseInt(filterd[0], 2, 64)
		}
	}

	return oxy * c02
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
