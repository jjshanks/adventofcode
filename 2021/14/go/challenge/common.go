package challenge

import (
	"regexp"
	"strings"
)

type Key struct {
	L byte
	R byte
}

func parse(instr string) (map[Key]byte, []byte) {
	lines := strings.Split(instr, "\n")
	template := []byte(lines[0])
	lines = lines[2:]
	pairs := make(map[Key]byte)
	r := regexp.MustCompile(`([A-Z]+) -> ([A-Z])`)
	for _, line := range lines {
		captures := r.FindStringSubmatch(line)
		key := Key{
			L: captures[1][0],
			R: captures[1][1],
		}
		pairs[key] = captures[2][0]

	}
	return pairs, template
}

func process(template []byte, pairs map[Key]byte, times int) int {
	p := make(map[byte]map[byte]byte)
	totals := make(map[byte]map[byte]uint64)
	// init pair and total maps
	for k, v := range pairs {
		if _, ok := p[k.L]; !ok {
			p[k.L] = make(map[byte]byte)
			totals[k.L] = make(map[byte]uint64)
		}
		p[k.L][k.R] = v
		totals[k.L][k.R] = 0
	}
	// seed initial values
	for i := 0; i < len(template)-1; i += 1 {
		l, r := template[i], template[i+1]
		c := p[l][r]
		totals[l][c] += 1
		totals[c][r] += 1
	}
	// iterate the rest
	for i := 1; i < times; i += 1 {
		deltas := make(map[Key]uint64)
		for l, ra := range totals {
			for r, n := range ra {
				c := p[l][r]
				deltas[Key{L: l, R: c}] += n
				deltas[Key{L: c, R: r}] += n
				// reset value as it gets replaced
				totals[l][r] = 0
			}
		}
		for k, n := range deltas {
			totals[k.L][k.R] += n
		}
	}
	// need to count left and right occurences to account for the last letter
	counts := make(map[byte]uint64)
	rightCounts := make(map[byte]uint64)
	for l, ra := range totals {
		for r, n := range ra {
			counts[l] += n
			rightCounts[r] += n
		}
	}
	for k, v := range rightCounts {
		if v > counts[k] {
			counts[k] = v
		}
	}
	min, max := uint64(0), uint64(0)
	for _, c := range counts {
		if min == 0 || min > c {
			min = c
		}
		if max < c {
			max = c
		}
	}
	return int(max - min)
}
