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
			`4
			10
			8
			8
			6
			`,
			"3",
		},
		{
			`3
			15
			15
			15
			`,
			"1",
		},
		{
			`7
			50
			30
			50
			100
			50
			80
			30
			`,
			"4",
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
