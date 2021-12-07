package main

import (
	"adventOfCode/helper"
	"os"
	"testing"
)

func TestDay07(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		r, _ := os.ReadFile("test.txt")
		s, _ := helper.InputToIntSlice(string(r), ",")

		got := part1(s)
		expected := 37

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Part2", func(t *testing.T) {
		r, _ := os.ReadFile("test.txt")
		s, _ := helper.InputToIntSlice(string(r), ",")

		got := part2(s)
		expected := 168

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
