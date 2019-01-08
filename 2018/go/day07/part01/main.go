package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("input")
	trip(err)
	defer f.Close() //nolint

	sc := bufio.NewScanner(f)

	ss := make(steps)

	for sc.Scan() {
		s, d := stepAndDep(sc.Text())
		ss[s] = append(ss[s], d)
	}
	trip(sc.Err())

	fmt.Println(ordered(ss))
}

type steps map[byte][]byte

func stepAndDep(s string) (byte, byte) {
	ws := strings.Split(s, " ")
	return ws[7][0], ws[1][0]
}

func ordered(ss steps) string {
	f := first(ss)
	undone := make([]byte, 0)
	out := make([]byte, 0)

	return string(apply(ss, f, undone, out))
}

func apply(ss steps, key byte, undone, out []byte) []byte {
	out = append(out, key)

	avs := make([]byte, 0)
	for av, pas := range ss {
		// if is not referrer, skip
		if !has(pas, key) {
			continue
		}

		//if every pa not done, skip
		if !hasAll(out, pas) {
			continue
		}

		avs = append(avs, av)
	}

	for _, un := range undone {
		if !has(avs, un) {
			avs = append(avs, un)
		}
	}

	alphabetize(avs)

	if len(avs) == 0 {
		return out
	}

	return apply(ss, avs[0], avs[1:], out)
}

func first(ss steps) byte {
	lu := make(map[byte]struct{})
	for _, bs := range ss {
		for _, b := range bs {
			lu[b] = struct{}{}
		}
	}

	for b := range lu {
		if _, ok := ss[b]; !ok {
			return b
		}
	}

	panic("cannot find first")
}

func has(bs []byte, b byte) bool {
	for _, v := range bs {
		if v == b {
			return true
		}
	}

	return false
}

func hasAll(bs, xs []byte) bool {
	for _, x := range xs {
		if !has(bs, x) {
			return false
		}
	}

	return true
}

type bytes []byte

func (s bytes) Len() int           { return len(s) }
func (s bytes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s bytes) Less(i, j int) bool { return s[i] < s[j] }

func alphabetize(bs []byte) {
	sort.Sort(bytes(bs))
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
