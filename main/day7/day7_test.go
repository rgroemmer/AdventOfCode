package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestBags(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		s := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
		var got int
		expected := 4
		for _, p := range strings.Split(string(s), "\n") {
			fmt.Println(createBags(p))
		}

		if !equal(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func equal(a, b int) bool {
	if a != b {
		return false
	}
	return true
}
