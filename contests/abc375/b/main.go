package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	scanInt, _ := createScanner(r)

	N := scanInt()

	X := make([]int, N)
	Y := make([]int, N)

	for i := 0; i < N; i++ {
		X[i], Y[i] = scanInt(), scanInt()
	}

	var sum float64

	sum += calcDist(0, 0, X[0], Y[0]) + calcDist(0, 0, X[N-1], Y[N-1])
	for i := 0; i < N-1; i++ {
		sum += calcDist(X[i], Y[i], X[i+1], Y[i+1])
	}

	fmt.Fprintln(w, sum)
}

func calcDist(a int, b int, c int, d int) float64 {
	return math.Sqrt(float64((a - c)*(a - c) + (b - d)*(b - d)))
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
