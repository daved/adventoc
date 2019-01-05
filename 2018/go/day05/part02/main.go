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

	p := makeProcessedStripped(chem)

	fmt.Println(len(p.short()))
}

type processed map[byte]string

func makeProcessedStripped(s string) processed {
	m := make(map[byte]string)

	for n := 65; n < 91; n++ {
		b := byte(n)
		m[b] = process(strip(s, b))
	}

	return m
}

func (p processed) short() string {
	low := int(1e+6)
	var s string

	for _, v := range p {
		if len(v) < low {
			low = len(v)
			s = v
		}
	}

	return s
}

func strip(s string, b byte) string {
	bs := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i]|0x20 == b|0x20 {
			continue
		}

		bs = append(bs, s[i])
	}

	return string(bs[:len(bs)])
}

func process(s string) string {
	bs := []byte(s)

	for i := 0; i < len(bs)-1; i++ {
		if !isSameLetterDiffCap(bs[i], bs[i+1]) {
			continue
		}

		bs = append(bs[:i], bs[i+2:]...)
		i -= 2
		if i < 0 {
			i = -1
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
