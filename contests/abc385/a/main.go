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

	// solve here
	A, B, C := scanInt(), scanInt(), scanInt()

	conditions := []bool{
		A == B && B == C,
		A + B == C,
		B + C == A,
		C + A == B,
	}

	for _, condition := range conditions {
		if condition {
			fmt.Fprintln(w, "Yes")
			return
		}
	}

	fmt.Fprintln(w, "No")
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
