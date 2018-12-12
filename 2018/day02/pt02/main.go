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

	var ids []string
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		ids = append(ids, sc.Text())
	}
	trip(sc.Err())

	for _, id := range ids {
		for _, ref := range ids {
			if asciiDiffs(ref, id) == 1 {
				asciiPrintCommon(ref, id)
				return
			}
		}
	}
}

func asciiDiffs(a, b string) int {
	var c int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			c++
		}
	}

	return c
}

func asciiPrintCommon(a, b string) {
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			fmt.Print(string(a[i]))
		}
	}

	fmt.Println()
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
