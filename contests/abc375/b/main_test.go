package main

import (
	"math"
	"strconv"
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
			1 2
			-1 0
			`,
			"6.06449510224597979401",
		},
		{
			`7
			-14142 13562
			-17320 50807
			-22360 67977
			24494 89742
			-26457 51311
			28284 27124
			31622 77660
			`,
			"384694.57587932075868509383",
		},
		{
			`5
			-100000 100000
			100000 -100000
			-100000 100000
			100000 -100000
			-100000 100000
			`,
			"1414213.56237309504880168872",
		},
	}

	for _, c := range cases {
		r := strings.NewReader(c.Input)
		var w strings.Builder

		solve(r, &w)

		got, err := strconv.ParseFloat(strings.Trim(w.String(), "\n "), 64)
		if err != nil {
			t.Fatalf(err.Error())
		}

		want, err := strconv.ParseFloat(c.Want, 64)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if math.Abs((want - got) / want) >= 1e-6 {
			t.Fatalf("want %v, got %v", want, got)
		}
	}
}
