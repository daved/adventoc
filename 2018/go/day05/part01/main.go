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
	var chem string

	for sc.Scan() {
		chem = sc.Text()
	}
	trip(sc.Err())

	res := process(chem)

	fmt.Println(len(res))
}

func process(s string) string {
	bs := []byte(s)

	for i := 0; i < len(bs)-1; i++ {
		if !isSameLetterDiffCap(bs[i], bs[i+1]) {
			continue
		}

		bs = append(bs[:i], bs[i+2:]...)
		i--
		if i > 0 {
			i -= 2
		}
	}

	return string(bs)
}

func isSameLetterDiffCap(a, b byte) bool {
	return a|0x20 == b|0x20 && a != b
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
