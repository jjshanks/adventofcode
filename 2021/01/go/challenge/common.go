package challenge

import (
	"strconv"
	"strings"
)

func parse(instr string) []int {
	v := make([]int, 0)
	for _, s := range strings.Split(instr, "\n") {
		i, _ := strconv.Atoi(s)
		v = append(v, i)
	}
	return v
}
