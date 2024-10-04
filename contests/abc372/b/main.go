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

func solve(r io.Reader, w io.Writer) []int {
	scanInt, _ := createScanner(r)

	M := scanInt()

	var A []int

	for i := 0; M > 0; i++ {
		rem := M % 3

		for rem > 0 {
			A = append(A, i)
			rem -= 1
		}

		M /= 3
	}

	fmt.Fprintln(w, len(A))

	for i, a := range(A) {
		if i < len(A) - 1 {
			fmt.Fprintf(w, "%v ", a)
		} else {
			fmt.Fprintf(w, "%v\n", a)
		}
	}

	return A
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
