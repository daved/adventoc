package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	trip(err)
	defer f.Close() //nolint

	sc := bufio.NewScanner(f)
	var cs coords

	for sc.Scan() {
		cs = append(cs, newCoord(sc.Text()))
	}
	trip(sc.Err())

	fmt.Println(largestAreaIn(500, cs))
}

func largestAreaIn(size int, cs coords) (id int, area int) {
	m := make(map[int]int)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			c := &coord{x: x, y: y}
			cl := closest(cs, c)

			if cl < 0 {
				continue
			}

			markInfinites(cs, cl, c)

			m[cl]++
		}
	}

	for k, v := range m {
		if !cs[k].inf && v > area {
			area = v
			id = k
		}
	}

	return id, area
}

func markInfinites(cs coords, id int, c *coord) {
	if id < 0 || id >= len(cs) {
		return
	}

	ref := cs[id]

	if ref.y == c.y && ref.x == c.x {
		return
	}

	diffy := c.y - ref.y
	diffx := c.x - ref.x
	ck := &coord{x: c.x + (1e+6 * diffx), y: c.y + (1e+6 * diffy)}

	if closest(cs, ck) == id {
		ref.inf = true
	}
}

func closest(refs coords, c *coord) int {
	var id int
	near := 1<<31 - 1
	var collision bool

	for i, ref := range refs {
		d := distance(ref, c)

		if d == near {
			collision = true
		}

		if d < near {
			collision = false
			near = d
			id = i
		}
	}

	if collision {
		return -1
	}

	return id
}

func distance(ref, c *coord) int {
	return abs(ref.y-c.y) + abs(ref.x-c.x)
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

type coord struct {
	x, y int
	inf  bool
}

func newCoord(s string) *coord {
	ss := strings.Split(s, ", ")

	x, err := strconv.Atoi(ss[0])
	trip(err)
	y, err := strconv.Atoi(ss[1])
	trip(err)

	return &coord{x: x, y: y}
}

type coords []*coord

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
