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

	for sc.Scan() {
		_ = sc.Text()
	}
	trip(sc.Err())

	fmt.Println("answer")
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
