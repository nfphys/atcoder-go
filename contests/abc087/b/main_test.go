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
	// add cases here
	cases := []Case{
		{
			`2
			2
			2
			100
			`,
			"2",
		},
		{
			`5
			1
			0
			150
			`,
			"0",
		},
		{
			`30
			40
			50
			6000
			`,
			"213",
		},
	}

	for _, c := range cases {
		r := strings.NewReader(c.Input)
		var w strings.Builder

		solve(r, &w)

		got := strings.Trim(w.String(), "\n ")

		if c.Want != got {
			t.Fatalf("want %s, got %s", c.Want, got)
		}
	}
}
