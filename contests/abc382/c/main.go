package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	r := os.Stdin

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	solve(r, w)
}

func solve(r io.Reader, w io.Writer) {
	scanInt, _ := createScanner(r)

	// solve here
	N, M := scanInt(), scanInt()

	A := make([]int, N)
	for i := range A {
		A[i] = scanInt()
	}

	B := make([]int, M)
	for i := range B {
		B[i] = scanInt()
	}

	maxB := 2 * 100000

	m := make(map[int]int, maxB)
	eatableMax := maxB
	for i, a := range A {
		if eatableMax >= a {
			for j := a; j <= eatableMax; j++ {
				m[j] = i
			}
			eatableMax = a-1
		}
	}

	for _, b := range B {
		if i, ok := m[b]; ok {
			fmt.Fprintln(w, i + 1)
		} else {
			fmt.Fprintln(w, -1)
		}
	}
}

func createScanner(r io.Reader) (func() int, func() string) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 1024*1024*100), 1024*1024*1000)
	sc.Split(bufio.ScanWords)

	scanInt := func() int {
		sc.Scan()
		i, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		return i
	}

	scanWord := func() string {
		sc.Scan()
		return sc.Text()
	}

	return scanInt, scanWord
}
