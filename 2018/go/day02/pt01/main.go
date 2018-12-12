package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input")
	trip(err)
	defer f.Close() //nolint

	sc := bufio.NewScanner(f)
	var dubs, trps int

	for sc.Scan() {
		d, t := twosAndThrees(sc.Text())
		dubs += d
		trps += t
	}
	trip(sc.Err())

	fmt.Println(dubs * trps)
}

func twosAndThrees(s string) (int, int) {
	m := make(map[rune]int)
	var d, t int

	for _, r := range s {
		m[r]++

		switch {
		case m[r] == 2:
			d++
		case m[r] == 3:
			d--
			t++
		case m[r] > 3:
			t--
		}
	}

	return oneOrNone(d > 0), oneOrNone(t > 0)
}

func oneOrNone(b bool) int {
	if b {
		return 1
	}
	return 0
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
