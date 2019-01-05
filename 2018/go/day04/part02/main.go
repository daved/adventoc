package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("input")
	trip(err)
	defer f.Close() //nolint

	obs := makeObservations(f)
	gsd := makeGuardsData(obs)

	fmt.Println(gsd.freqIDMultPeak())
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}

func parseTime(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04", s)
	trip(err)

	return t
}

func parseID(s string) int {
	if len(s) < 2 || s[0] != '#' {
		return 0
	}

	n, err := strconv.Atoi(s[1:])
	trip(err)

	return n
}

const (
	sleepTkn = "falls"
)

type observation struct {
	id  int
	t   time.Time
	tkn string
}

func makeObservation(msg string) observation {
	chunks := strings.Split(msg, "]")
	words := strings.Split(strings.TrimSpace(chunks[1]), " ")

	return observation{
		id:  parseID(words[1]),
		t:   parseTime(chunks[0][1:]),
		tkn: words[0],
	}
}

type observations []observation

func makeObservations(r io.Reader) observations {
	sc := bufio.NewScanner(r)
	obs := make(observations, 0)

	for sc.Scan() {
		obs = append(obs, makeObservation(sc.Text()))
	}
	trip(sc.Err())

	sort.Sort(obs)
	obs = updateIDs(obs)

	return obs
}

func (obs observations) Len() int           { return len(obs) }
func (obs observations) Swap(i, j int)      { obs[i], obs[j] = obs[j], obs[i] }
func (obs observations) Less(i, j int) bool { return obs[i].t.Before(obs[j].t) }

// must be sorted by time
func updateIDs(obs observations) observations {
	cid := 0

	for i, ob := range obs {
		if ob.id != 0 {
			cid = ob.id
			continue
		}

		obs[i].id = cid
	}

	return obs
}

type guardData struct {
	peak  int // minute of hour
	freq  int
	total int // minutes
	obs   observations
}

type guardsData map[int]*guardData

func makeGuardsData(obs observations) guardsData {
	m := make(map[int]*guardData)

	for _, ob := range obs {
		gd, ok := m[ob.id]
		if !ok {
			gd = &guardData{}
			m[ob.id] = gd
		}

		gd.obs = append(gd.obs, ob)
	}

	for id, gd := range m {
		m[id].peak, m[id].freq, m[id].total = peakFreqTotal(gd.obs)
	}

	return m
}

func (gsd guardsData) topIDMultPeak() int {
	var top, id, pk int

	for gid, gd := range gsd {
		if gd.total > top {
			top = gd.total
			id = gid
			pk = gd.peak
		}
	}

	return id * pk
}

func (gsd guardsData) freqIDMultPeak() int {
	var top, id, pk int

	for gid, gd := range gsd {
		if gd.freq > top {
			top = gd.freq
			id = gid
			pk = gd.peak
		}
	}

	return id * pk
}

func applySpan(a [60]int, index, ct int) [60]int {
	if index+ct > len(a) {
		panic("out of span")
	}

	for i := index; i < index+ct; i++ {
		a[i]++
	}

	return a
}

func findPeak(a [60]int) int {
	var min, p int

	for m, n := range a {
		if n > p {
			p = n
			min = m
		}
	}

	return min
}

func peakFreqTotal(obs observations) (int, int, int) {
	var total int
	var tally [60]int

	for i, ob := range obs {
		if ob.tkn != sleepTkn {
			continue
		}

		j := ob.t.Minute()

		if i+1 >= len(obs) {
			panic("this one never wakes")
		}

		ct := obs[i+1].t.Minute() - j

		if ct < 0 {
			tct := len(tally) - j
			tally = applySpan(tally, j, tct)
			total += tct

			ct = j
			j = 0
		}

		tally = applySpan(tally, j, ct)
		total += ct
	}

	pk := findPeak(tally)

	return pk, tally[pk], total
}
