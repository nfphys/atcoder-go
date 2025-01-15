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

type Snake struct {
  T int
  L int
}

func solve(r io.Reader, w io.Writer) {
	scanInt, _ := createScanner(r)

	// solve here
  N, D := scanInt(), scanInt()

  snakes := make([]Snake, N)
  for i := 0; i < N; i++ {
    T, L := scanInt(), scanInt()
    snakes[i] = Snake{ T, L }
  }

  for k := 1; k <= D; k++ {
    max := 0
    for _, snake := range snakes {
      weight := snake.T * (snake.L + k)
      if weight > max {
        max = weight
      }
    }
    fmt.Fprintln(w, max)
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
