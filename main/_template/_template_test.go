package main

import (
	"testing"
)

func TestBags(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		s := "L.LL.LL.LL,LLLLLLL.LL,L.L.L..L..,LLLL.LL.LL,L.LL.LL.LL,L.LLLLL.LL,..L.L.....,LLLLLLLLLL,L.LLLLLL.L,L.LLLLL.LL"

		got := part1(s)
		expected := 37

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
