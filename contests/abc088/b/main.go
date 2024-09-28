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
	a := make([]int, N)

	for i := 0; i < N; i++ {
		a[i] = scanInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	alice := 0
	bob := 0

	for i := 0; i < N; i++ {
		if i % 2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}

	fmt.Fprintln(w, alice - bob)
}
