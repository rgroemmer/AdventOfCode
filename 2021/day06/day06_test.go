package main

import (
	"adventOfCode/helper"
	"os"
	"testing"
)

func TestDay06(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		r, _ := os.ReadFile("test.txt")
		s, _ := helper.InputToIntSlice(string(r), ",")

		got := part2(s, 80)
		expected := 5934

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		r, _ := os.ReadFile("input.txt")
		s, _ := helper.InputToIntSlice(string(r), ",")

		got := part2(s, 256)
		expected := 1740449478328

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

}

func BenchmarkPart1(b *testing.B) {
	r, _ := os.ReadFile("input.txt")
	s, _ := helper.InputToIntSlice(string(r), ",")
	for n := 0; n < b.N; n++ {
		_ = part2(s, 256)
	}
}

func BenchmarkPart2(b *testing.B) {
	r, _ := os.ReadFile("input.txt")
	s, _ := helper.InputToIntSlice(string(r), ",")
	for n := 0; n < b.N; n++ {
		_ = part2(s, 256)
	}
}
