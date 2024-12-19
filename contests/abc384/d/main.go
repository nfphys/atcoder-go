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

	N, S := scanInt(), int64(scanInt())

	A := make([]int64, N)
	for i := 0; i < N; i++ {
		A[i] = int64(scanInt())
	}

	sumA := int64(0)
	for _, a := range A {
		sumA += a
	}

	S = S % sumA

	L := 0
	R := 0

	T := A[0]

	for L < 2*N && R < 2*N {
		switch {
		case T == S:
			fmt.Fprintln(w, "Yes")
			return
		case T < S:
			R += 1
			T += A[R % N]
		case T > S:
			T -= A[L % N]
			L += 1
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
