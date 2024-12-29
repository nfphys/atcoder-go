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

	// solve here
	var H, W, X, Y int
	H, W, X, Y = scanInt(), scanInt(), scanInt() - 1, scanInt() - 1

	S := make([][]rune, H)
	for i := 0; i < H; i++ {
		S[i] = []rune(scanWord())
	}

	T := scanWord()

	visited := make([][]bool, H)
	for i := 0; i < H; i++ {
		visited[i] = make([]bool, W)
	}

	x, y := X, Y
	visited[x][y] = true

	for _, t := range T {
		switch t {
		case 'U':
			if S[x-1][y] != '#' {
				x = x-1
			}
		case 'D':
			if S[x+1][y] != '#' {
				x = x+1
			}
		case 'L':
			if S[x][y-1] != '#' {
				y = y-1
			}
		case 'R':
			if S[x][y+1] != '#' {
				y = y+1
			}
		}
		visited[x][y] = true
	}

	count := 0

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == '@' && visited[i][j] {
				count += 1
			}
		}
	}

	fmt.Fprintln(w, x+1, y+1, count)
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
