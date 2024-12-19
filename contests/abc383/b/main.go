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
	scanInt, scanWord := createScanner(r)

	var H, W, D int
	H, W, D = scanInt(), scanInt(), scanInt()

	S := make([][]rune, H)
	for h := 0; h < H; h++ {
		S[h] = []rune(scanWord())
	}

	max := 0

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == '#' {
				continue
			}

			for k := 0; k < H; k++ {
				for l := 0; l < W; l++ {
					if S[k][l] == '#' {
						continue
					}

					count := 0

					for m := 0; m < H; m++ {
						for n := 0; n < W; n++ {
							if S[m][n] == '#' {
								continue
							}

							flag := false

							if abs(i-m) + abs(j-n) <= D {
								flag = true
							}

							if abs(k-m) + abs(l-n) <= D {
								flag = true
							}

							if flag {
								count += 1
							}
						}
					}

					if count > max {
						max = count
					}
				}
			}
		}	
	}

	fmt.Fprintln(w, max)
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
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
