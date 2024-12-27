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
	r := os.Stdin

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	solve(r, w)
}

func solve(r io.Reader, w io.Writer) {
	scanInt, _ := createScanner(r)

	// solve here
	N, M := scanInt(), scanInt()

	results := []string{}

	var f func([]int)
	f = func(A []int) {
		if len(A) > N {
			return
		}

		if len(A) == N {
			s := make([]string, N)
			for i, a := range A {
				s[i] = strconv.Itoa(a)
			}
			results = append(results, strings.Join(s, " "))
			return
		}

		var min int
		if len(A) == 0 {
			min = 1
		} else {
			min = A[len(A)-1] + 10
		}

		max := M - (N - len(A) - 1) * 10

		if min > max {
			return
		}

		for i := min; i <= max; i++ {
			A = append(A, i)
			f(A)
			A = A[:len(A)-1]
		}

	}

	f([]int{})

	fmt.Fprintln(w, len(results))
	for _, r := range results {
		fmt.Fprintln(w, r)
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
