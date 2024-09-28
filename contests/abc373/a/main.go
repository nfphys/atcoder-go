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
	sc.Split(bufio.ScanWords)

	scanWord := func () string {
		sc.Scan()
		return sc.Text()
	}

	count := 0

	for i := 1; i <= 12; i++ {
		if len(scanWord()) == i {
			count += 1
		}
	}

	fmt.Fprintln(w, count)
}
