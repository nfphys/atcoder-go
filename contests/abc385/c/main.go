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
  N := scanInt()

  H := make([]int, N)
  for i := range H {
    H[i] = scanInt()
  }

  max := 1

  for interval := 1; interval <= N-1; interval++ {
    for start := 0; start < interval; start++ {
      count := 0
      height := 0

      for i := start; i < N; i += interval {
        if H[i] != height {
          count = 0
          height = H[i]
        }

        count += 1

        if count > max {
          max = count
        }
      }
    }
  }

  fmt.Fprintln(w, max)
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
