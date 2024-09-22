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

	n := scanInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	count := 0
	for {
		for i := 0; i < n; i++  {
			if a[i] % 2 != 0 {
				fmt.Fprintln(w, count)
				return
			}

			a[i] /= 2
		}

		count += 1
	}
}
