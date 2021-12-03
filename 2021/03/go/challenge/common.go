package challenge

import "strings"

func parse(instr string) [][]bool {
	r := make([][]bool, 0)
	for _, line := range strings.Split(instr, "\n") {
		inner := make([]bool, 0)
		for _, b := range []byte(line) {
			inner = append(inner, b == byte(49))
		}
		r = append(r, inner)
	}
	return r
}

func convert(b []bool) (result int) {
	for idx, _ := range b {
		shift := len(b) - idx - 1
		if b[idx] {
			result += 1 << shift
		}
	}
	return
}
