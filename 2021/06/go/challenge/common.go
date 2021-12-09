package challenge

import (
	"strings"
)

func parse(instr string) []int {
	matches := make(map[string]int)
	for _, s := range strings.Split(instr, ",") {
		matches[s] += 1
	}
	a := make([]int, 7)
	for i, k := range []string{"0", "1", "2", "3", "4", "5", "6"} {
		a[i] = matches[k]
	}
	return a
}
