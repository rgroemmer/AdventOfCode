package main

import (
	helper "../helper"
	"os"
	"strings"
)

func countAnyone(group string) int {
	var result string
	for _, char := range group {
		if !strings.Contains(result, string(char)) && char != ',' {
			result += string(char)
		}
	}
	return len(result)
}

func isInOther(char string, others []string) bool {
	var count int
	for _, answers := range others {
		if strings.Contains(answers, char) {
			count++
		}
	}
	return count == len(others)
}

func countEveryone(group string) int {
	first := strings.Split(group, ",")[0]
	others := strings.Split(group, ",")[1:]
	var count int
	for _, char := range first {
		if isInOther(string(char), others) {
			count++
		}
	}
	return count
}

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day6/input.txt")
	if err != nil {
		println("Error: ", err)
	}
	groups := strings.Split(helper.ReadeFile(r), "|")
	var magicNumber int
	var magicNumber2 int

	for _, group := range groups {
		magicNumber += countAnyone(group)
		magicNumber2 += countEveryone(group)
	}
	println("MagicNumber: ", magicNumber)
	println("MagicNumber2: ", magicNumber2)
}
