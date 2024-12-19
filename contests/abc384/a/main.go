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
	scanInt, scanWord := createScanner(r)

	_, c1, c2 := scanInt(), rune(scanWord()[0]), rune(scanWord()[0])

	S := []rune(scanWord())

	T := make([]rune, len(S))

	for i, c := range S {
		if c != c1 {
			T[i] = c2
		} else {
			T[i] = c1
		}
	}

	fmt.Fprintln(w, string(T))
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
