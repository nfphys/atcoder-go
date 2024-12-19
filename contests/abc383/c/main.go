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

const INF int = 10000000000

func solve(r io.Reader, w io.Writer) {
	scanInt, scanWord := createScanner(r)

	var H, W, D int
	H, W, D = scanInt(), scanInt(), scanInt()

	S := make([][]rune, H)
	for i := 0; i < H; i++ {
		S[i] = []rune(scanWord())
	}

	q := make([][]int, 0)

	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = INF
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == 'H' {
				q = append(q, []int{i, j})
				dist[i][j] = 0
			}
		}	
	}

	search := func(i, j, d int) {
		if i < 0 || i >= H {
			return
		}

		if j < 0 || j >= W {
			return
		}

		if d > D {
			return
		}

		if S[i][j] == '#' {
			return
		}

		if dist[i][j] < d {
			return
		}

		dist[i][j] = d
		q = append(q, []int{i, j})
	}

	for len(q) > 0 {
		i := q[0][0]
		j := q[0][1]
		q = q[1:]

		d := dist[i][j]
		search(i+1, j, d+1)
		search(i-1, j, d+1)
		search(i, j+1, d+1)
		search(i, j-1, d+1)
	}

	count := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if dist[i][j] < INF {
				count += 1
			}
		}
	}

	fmt.Fprintln(w, count)
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
