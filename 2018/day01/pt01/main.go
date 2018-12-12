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

	sc := bufio.NewScanner(f)
	var a int

	for sc.Scan() {
		c, err := strconv.Atoi(sc.Text())
		trip(err)
		a += c
	}
	trip(sc.Err())

	fmt.Println(a)
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
