package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	reader := os.Stdin

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solve(reader, writer)	
}

func solve(r io.Reader, w io.Writer) {
	scanInt, scanWord := createScanner(r)

	N := scanInt()
	S := scanWord()

	count := 0
	for i := 0; i < N-2; i++ {
		if S[i:i+3] == "#.#" {
			count++
		}
	}

	fmt.Fprintln(w, count)
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
