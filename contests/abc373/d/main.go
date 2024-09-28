package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

type Edge struct {
	From int
	To int
	Weight int
}

func (e Edge) Reverse() Edge {
	return Edge{
		From: e.To,
		To: e.From,
		Weight: -e.Weight,
	}
}

func solve(r io.Reader, w io.Writer) ([]Edge, []int) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	scanInt := func() int {
		sc.Scan()
		i, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		return i
	}

	N := scanInt()
	M := scanInt()

	edges := make([]Edge, M)

	for i := 0; i < M; i++ {
		edges[i] = Edge{
			From: scanInt() - 1,
			To: scanInt() - 1,
			Weight: scanInt(),
		}
	}

	g := make([][]Edge, N)

	for i := 0; i < M; i++ {
		e := edges[i]
		from := e.From
		to := e.To

		g[from] = append(g[from], e)
		g[to] = append(g[to], e.Reverse())
	}

	seen := make([]bool, N)
	X := make([]int, N)

	var search func(int, int)
	search = func (s int, x int) {
		if seen[s] {
			return
		}

		seen[s] = true
		X[s] = x

		for _, e := range g[s] {
			search(e.To, x + e.Weight)
		}
	}

	for s := 0; s < N; s++ {
		search(s, 0)
	}

	strs := make([]string, len(X))
	for i, x := range X {
		strs[i] = strconv.Itoa(x)
	}

	fmt.Fprintln(w, strings.Join(strs, " "))

	return edges, X
}
