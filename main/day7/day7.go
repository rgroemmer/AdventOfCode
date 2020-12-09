package main

import (
	"os"
	"strconv"
	"strings"
)

type bag struct {
	color   string
	contain map[string]int
}

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day7/input.txt")
	if err != nil {
		println("Error: ", err)
	}
	input := helper.ReadeFileByLine(r)

	createBags("posh black bags contain 3 dark lavender bags, 3 mirrored coral bags, 1 dotted chartreuse bag.")
}

func createBags(s string) bag {
	bagRules := strings.Split(s, " ")
	if bagRules[4] == "no" {
		return bag{
			color: bagRules[0] + " " + bagRules[1],
		}
	}
	if len(bagRules) == 8 {
		count, _ := strconv.Atoi(bagRules[4])
		return bag{
			color: bagRules[0] + " " + bagRules[1],
			contain: map[string]int{
				bagRules[5] + " " + bagRules[6]: count,
			},
		}
	}
	if len(bagRules) > 8 {
		contains := strings.Split(strings.Split(s, "contain ")[1], ", ")
		containBags := make(map[string]int)
		for _, contain := range contains {
			split := strings.Split(contain, " ")
			color := split[1] + " " + split[2]
			count := split[0]
			containBags[color], _ = strconv.Atoi(count)
		}
		return bag{
			color:   bagRules[0] + " " + bagRules[1],
			contain: containBags,
		}
	}
	return bag{}
}
