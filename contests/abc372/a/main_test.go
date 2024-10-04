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
		{ ".v.", "v" },
		{ "chokudai", "chokudai" },
		{ "...", "" },
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
