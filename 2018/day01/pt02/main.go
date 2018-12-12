package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	trip(err)
	defer f.Close() //nolint

	var a int
	m := make(map[int]struct{})
	var r *int

	for {
		sc := bufio.NewScanner(f)

		for sc.Scan() {
			c, err := strconv.Atoi(sc.Text())
			trip(err)
			a += c

			if _, ok := m[a]; ok {
				r = &a
				break
			}
			m[a] = struct{}{}
		}
		trip(sc.Err())

		if r != nil {
			break
		}
		_, _ = f.Seek(0, 0) //nolint
	}

	fmt.Println(*r)
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
