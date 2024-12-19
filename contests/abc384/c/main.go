package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Player struct{
	name string
	point int
}

type Players []Player

func (ps Players) Len() int {
	return len(ps)
}

func (ps Players) Less(i, j int) bool {
	if ps[i].point != ps[j].point {
		return ps[i].point < ps[j].point
	} else {
		return ps[i].name > ps[j].name
	}
}

func (ps Players) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	r := os.Stdin

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	solve(r, w)
}

func solve(r io.Reader, w io.Writer) {
	scanInt, _ := createScanner(r)

	a, b, c, d, e := scanInt(), scanInt(), scanInt(), scanInt(), scanInt()

	points := map[rune]int{
		'A': a,
		'B': b,
		'C': c,
		'D': d,
		'E': e,
	}
	
	players := Players{}

	var f func(name string, point int, c rune)
	f = func (name string, point int, c rune) {
		if c > 'E' {
			if name != "" {
				players = append(players, Player{ name, point })
			}

			return
		}

		f(name + string(c), point + points[c], c + 1)
		f(name, point, c + 1)
	}

	f("", 0, 'A')

	sort.Stable(sort.Reverse(players))

	for _, p := range players {
		fmt.Fprintln(w, p.name)
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
