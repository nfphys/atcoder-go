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

	scanWord := func() string {
		sc.Scan()
		return sc.Text()
	}

	s := scanWord()

	// alphabet to coordinate
	m := make(map[rune]int, 30)

	for i, c := range s {
		x := i + 1
		m[c] = x
	}

	sum := 0
	xBefore := m['A']

	for i := 0; i < 26; i++ {
		c := rune('A' + i)
		sum += abs(m[c] - xBefore)
		xBefore = m[c]
	}

	fmt.Fprintln(w, sum)
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}
