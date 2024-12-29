package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	r := os.Stdin

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	solve(r, w)
}

type Point struct {
  X int
  Y int
}

func solve(r io.Reader, w io.Writer) {
	scanInt, scanWord := createScanner(r)

	// solve here
  N, M, Sx, Sy := scanInt(), scanInt(), scanInt(), scanInt()

  X := make([]int, N)
  Y := make([]int, N)
  for i := range X {
    X[i], Y[i] = scanInt(), scanInt()
  }

  D := make([]rune, M)
  C := make([]int, M)
  for i := range D {
    D[i], C[i] = []rune(scanWord())[0], scanInt()
  }

  xToYTree := make(map[int]*redblacktree.Tree)
  yToXTree := make(map[int]*redblacktree.Tree)

  for i := range X {
    x, y := X[i], Y[i]

    if tree, ok := xToYTree[x]; ok {
      tree.Put(y, nil)
    } else {
      tree = redblacktree.NewWithIntComparator()
      tree.Put(y, nil)
      xToYTree[x] = tree
    }

    if tree, ok := yToXTree[y]; ok {
      tree.Put(x, nil)
    } else {
      tree = redblacktree.NewWithIntComparator()
      tree.Put(x, nil)
      yToXTree[y] = tree
    }
  }

  currentX := Sx
  currentY := Sy

  count := 0

  for i := range D {
    d, c := D[i], C[i]

    houses := []Point{}

    switch d {
    case 'U':
      if yTree, ok := xToYTree[currentX]; ok {
        if node, found := yTree.Ceiling(currentY); found {
          iter := yTree.IteratorAt(node)
          iter.Prev()
          for iter.Next() {
            key := (iter.Key()).(int)
            if key > currentY + c {
              break
            }

            houses = append(houses, Point{ currentX, key })
          }
        }
      }

      currentY += c
    case 'D':
      if yTree, ok := xToYTree[currentX]; ok {
        if node, found := yTree.Ceiling(currentY - c); found {
          iter := yTree.IteratorAt(node)
          iter.Prev()
          for iter.Next() {
            key := (iter.Key()).(int)
            if key > currentY {
              break
            }

            houses = append(houses, Point{ currentX, key })
          }
        }
      }

      currentY -= c
    case 'L':
      if xTree, ok := yToXTree[currentY]; ok {
        if node, found := xTree.Ceiling(currentX - c); found {
          iter := xTree.IteratorAt(node)
          iter.Prev()
          for iter.Next() {
            key := (iter.Key()).(int)
            if key > currentX {
              break
            }

            houses = append(houses, Point{ key, currentY })
          }
        }
      }

      currentX -= c
    case 'R':
      if xTree, ok := yToXTree[currentY]; ok {
        if node, found := xTree.Ceiling(currentX); found {
          iter := xTree.IteratorAt(node)
          iter.Prev()
          for iter.Next() {
            key := (iter.Key()).(int)
            if key > currentX + c {
              break
            }

            houses = append(houses, Point{ key, currentY })
          }
        }
      }

      currentX += c
    }

    for _, house := range houses {
      if yTree, ok := xToYTree[house.X]; ok {
        yTree.Remove(house.Y)
      }

      if xTree, ok := yToXTree[house.Y]; ok {
        xTree.Remove(house.X)
      }

      count++
    }
  }

  fmt.Fprintln(w, currentX, currentY, count)
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
