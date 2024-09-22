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

	A := scanInt()
	B := scanInt()
	C := scanInt()
	X := scanInt()

	count := 0

	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			rem := (X - 500*a - 100*b)

			if rem < 0 || rem % 50 != 0 || rem / 50 > C {
				continue
			}

			count += 1
		}
	}

	fmt.Fprintln(w, count)
}
