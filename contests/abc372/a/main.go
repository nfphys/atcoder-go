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
	_, scanWord := createScanner(r)

	S := scanWord()

	for _, c := range S {
		if c != '.' {
			fmt.Fprint(w, string(c))
		}
	}
	fmt.Fprint(w, "\n")
}

func createScanner(r io.Reader) (func() int, func() string) {
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

	scanWord := func() string {
		sc.Scan()
		return sc.Text()
	}

	return scanInt, scanWord
}
