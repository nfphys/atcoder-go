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
			Input: `3 3
							1 2 2
							3 2 3
							1 3 -1
							`,
		},
		{
			Input: `4 2
							2 1 5
							3 4 -3
							`,
		},
		{
			Input: `5 7
							2 1 18169343
							3 1 307110901
							4 1 130955934
							2 3 -288941558
							2 5 96267410
							5 3 -385208968
							4 3 -176154967
							`,
		},
	}

	for _, c := range cases {
		r := strings.NewReader(c.Input)
		var w strings.Builder

		edges, X := solve(r, &w)

		for _, e := range edges {
			if X[e.To] - X[e.From] != e.Weight {
				t.Fatalf("condition not satisfied.")
			}
		}
	}
}
