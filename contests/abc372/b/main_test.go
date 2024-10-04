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
		{ Input: "6" },
		{ Input: "100" },
		{ Input: "59048" },
	}

	for _, c := range cases {
		r := strings.NewReader(c.Input)
		var w strings.Builder

		M, _ := strconv.Atoi(c.Input)

		A := solve(r, &w)

		sum := 0
		for _, a := range(A) {
			sum += int(math.Pow(3, float64(a)))
		}

		if sum != M {
			t.Fatal("condition not satisfied.")
		}
	}
}
