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

	N := scanInt()
	A := scanInt()
	B := scanInt()

	sum := 0

	for i := 1; i <= N; i++ {
		s := sumDigits(i)
		if A <= s && s <= B {
			sum += i
		}
	}

	fmt.Fprintln(w, sum)
}

func sumDigits(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}
