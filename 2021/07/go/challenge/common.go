package challenge

import (
	"strconv"
	"strings"
)

func parse(instr string) []int {
	parts := strings.Split(instr, ",")
	nums := make([]int, len(parts))
	for idx, s := range parts {
		i, _ := strconv.Atoi(s)
		nums[idx] = i
	}
	return nums
}
