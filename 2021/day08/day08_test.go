package main

import (
	"adventOfCode/helper"
	"os"
	"testing"
)

func TestDay08(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		r, _ := os.ReadFile("test.txt")
		s, _ := helper.InputToStringSlice(string(r))

		got := part1(s)
		expected := 26

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Part2", func(t *testing.T) {
		r, _ := os.ReadFile("test.txt")
		s, _ := helper.InputToStringSlice(string(r))

		got := part2(s)
		expected := 61229

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
