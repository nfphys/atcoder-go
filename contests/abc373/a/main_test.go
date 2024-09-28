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
			`january
			february
			march
			april
			may
			june
			july
			august
			september
			october
			november
			december
			`,
			"1",
		},
		{
			`ve
			inrtfa
			npccxva
			djiq
			lmbkktngaovl
			mlfiv
			fmbvcmuxuwggfq
			qgmtwxmb
			jii
			ts
			bfxrvs
			eqvy
			`,
			"2",
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
