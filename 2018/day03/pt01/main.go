package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		size = 1000
	)

	f, err := os.Open("input")
	trip(err)
	defer f.Close() //nolint

	sc := bufio.NewScanner(f)
	fab := makeFabric(size, size)

	for sc.Scan() {
		c := makeClaim(sc.Text())
		fab.apply(c)
	}
	trip(sc.Err())

	fmt.Println(fab.over())
}

type claim struct {
	id   int
	x, y int
	w, h int
}

func makeClaim(s string) claim {
	ss := strings.Split(s, "@")
	a := strings.TrimSpace(ss[0][1:])

	ss = strings.Split(ss[1], ":")
	b := strings.TrimSpace(ss[0])
	c := strings.TrimSpace(ss[1])

	id, err := strconv.Atoi(a)
	trip(err)

	ss = strings.Split(b, ",")
	x, err := strconv.Atoi(strings.TrimSpace(ss[0]))
	trip(err)
	y, err := strconv.Atoi(strings.TrimSpace(ss[1]))
	trip(err)

	ss = strings.Split(c, "x")
	w, err := strconv.Atoi(strings.TrimSpace(ss[0]))
	trip(err)
	h, err := strconv.Atoi(strings.TrimSpace(ss[1]))
	trip(err)

	return claim{
		id: id,
		x:  x,
		y:  y,
		w:  w,
		h:  h,
	}
}

type fabric [][]int

func makeFabric(x, y int) fabric {
	f := make([][]int, x)

	for i := range f {
		f[i] = make([]int, y)
	}

	return f
}

func (f fabric) apply(c claim) {
	for x := c.x; x < len(f) && x < c.x+c.w; x++ {
		for y := c.y; y < len(f[x]) && y < c.y+c.h; y++ {
			f[x][y]++
		}
	}
}

func (f fabric) over() int {
	var a int

	for i := range f {
		for _, v := range f[i] {
			if v > 1 {
				a++
			}
		}
	}

	return a
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
