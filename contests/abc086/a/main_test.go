package main

import (
	"strings"
	"testing"
)

type Case struct {
	Input string
	Want string
}

func TestSolve(t *testing.T) {
	cases := []Case{
		{"3 4", "Even\n"},
		{"1 21", "Odd\n"},
	}

	for _, c := range cases {
		r := strings.NewReader(c.Input)
		var w strings.Builder

		solve(r, &w)

		if w.String() != c.Want {
			t.Fatalf("want %s, got %s", c.Want, w.String())
		}
	}
}
