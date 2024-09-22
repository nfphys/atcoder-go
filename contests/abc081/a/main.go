package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)

	sc.Scan()
	s := sc.Text()

	count := 0
	for _, c := range s {
		if c == '1' {
			count += 1
		}
	}

	fmt.Fprintln(w, count)
}
