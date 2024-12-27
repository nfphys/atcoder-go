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

	// solve here
	N, D := scanInt(), scanInt()
	S := []rune(scanWord())

	result := make([]rune, N)
	
	count := 0

	for i := N-1; i >= 0; i-- {
		if count < D && S[i] == '@' {
			result[i] = '.'
			count += 1
		} else {
			result[i] = S[i]
		}
	}

	fmt.Fprintln(w, string(result))
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
