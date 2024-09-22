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
			`3
			8 12 40`,
			"2",
		},
		{
			`4
			5 6 8 10`,
			"0",
		},
		{
			`6
			382253568 723152896 37802240 379425024 404894720 471526144`,
			"8",
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
