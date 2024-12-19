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

	N := scanInt()

	adds := make(map[int]int, N)

	var end int

	for i := 0; i < N; i++ {
		t, v := scanInt(), scanInt()
		adds[t] = v

		if i == N-1 {
			end = t
		}
	}

	result := 0

	for t := 1; t <= end; t++ {
		if result > 0 {
			result -= 1
		}

		if v, ok := adds[t]; ok {
			result += v
		}
	}

	fmt.Fprintln(w, result)
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
