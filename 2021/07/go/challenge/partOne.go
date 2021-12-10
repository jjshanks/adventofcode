package challenge

import (
	"sort"
)

func PartOne(instr string) int {
	inputSlice := parse(instr)
	sort.Ints(inputSlice)
	medianIdx := len(inputSlice) / 2
	median := inputSlice[medianIdx]
	if medianIdx%2 == 0 {
		median = (inputSlice[medianIdx-1] + inputSlice[medianIdx]) / 2
	}
	total := 0
	for _, v := range inputSlice {
		abs := median - v
		if abs < 0 {
			abs = -abs
		}
		total += abs
	}
	return total
}
