package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
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

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	B := make([]int, N)
	for i := 0; i < N; i++ {
		B[i] = scanInt()
	}

	// sort descending
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	sort.Sort(sort.Reverse(sort.IntSlice(B)))

	fmt.Fprintln(w, A[0] + B[0])
}
