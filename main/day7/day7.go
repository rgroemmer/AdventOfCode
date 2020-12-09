package main

import "strings"

type bag struct {
	color   string
	contain map[int]string
}

func createBags(s string) bag {
	split := strings.Split(s, " ")
	return bag{
		color: split[0] + " " + split[1],
	}
}

func main() {
	lol := createBags("pale chartreuse bags contain 3 faded orange bags.")

	println(lol.color)
}
