package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dodontRegex := regexp.MustCompile(`don't\(\).*?(?:do\(\))`)
	mulRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	numsRegex := regexp.MustCompile(`\d{1,3}`)

	acc := &Accumulation{}

	singleLineInputData := strings.ReplaceAll(inputData, "\n", "")
	validInputData := dodontRegex.ReplaceAllString(singleLineInputData, "")

	mulMatches := mulRegex.FindAllString(validInputData, -1)
	for _, mulMatch := range mulMatches {
		numsMatches := numsRegex.FindAllString(mulMatch, -1)
		acc.add(numsMatches)
	}

	if acc.err != nil {
		fmt.Fprintln(os.Stderr, acc.err)
		os.Exit(1)
	}

	fmt.Println(acc.total)
}

type Accumulation struct {
	total int
	err   error
}

func (a *Accumulation) add(nums []string) {
	if a.err != nil {
		return
	}

	numsLen := len(nums)
	if numsLen != 2 {
		a.err = fmt.Errorf("nums matches length not 2: len = %d", numsLen)
		return
	}

	atoi := func(num string) int {
		if a.err != nil {
			return 0
		}

		n, err := strconv.Atoi(num)
		if err != nil {
			a.err = fmt.Errorf("atoi '%s,%s': %w", nums[0], nums[1], err)
			return 0
		}
		return n
	}

	a.total += atoi(nums[0]) * atoi(nums[1])
}
