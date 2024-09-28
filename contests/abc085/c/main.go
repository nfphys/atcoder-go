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

	N, Y := scanInt(), scanInt()

	for x := 0; x <= N; x++ {
		for y := 0; y <= N; y++ {
			z := N - x - y
			if z < 0 {
				continue
			}

			if 10000*x + 5000*y + 1000*z == Y {
				fmt.Fprintf(w, "%v %v %v", x, y, z)
				return
			}
		}
	}

	fmt.Fprintln(w, "-1 -1 -1")
}
