package main

import (
	"adventOfCode/helper"
	"os"
	"testing"
)

func TestDay05(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		r, _ := os.ReadFile("test.txt")
		s, _ := helper.InputToStringSlice(string(r))

		got := part1(s, 10)
		expected := 5

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Part2", func(t *testing.T) {
		r, _ := os.ReadFile("test.txt")
		s, _ := helper.InputToStringSlice(string(r))

		got := part2(s, 10)
		expected := 12

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
