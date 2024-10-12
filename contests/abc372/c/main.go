package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	scanInt, scanWord := createScanner(r)

	N, Q := scanInt(), scanInt()

	S := []rune(scanWord())

	count := 0
	for i := 0; i < N-2; i++ {
		if S[i] == 'A' && S[i+1] == 'B' && S[i+2] == 'C' {
			count++
		}
	}

	for i := 0; i < Q; i++ {
		X, C := scanInt()-1, rune(scanWord()[0])

		for j := X-2; j <= X; j++ {
			if !(0 <= j && j < N-2) {
				continue
			}

			if S[j] == 'A' && S[j+1] == 'B' && S[j+2] == 'C' {
				count--
			}
		}

		S[X] = C

		for j := X-2; j <= X; j++ {
			if !(0 <= j && j < N-2) {
				continue
			}

			if S[j] == 'A' && S[j+1] == 'B' && S[j+2] == 'C' {
				count++
			}
		}

		fmt.Fprintln(w, count)
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
