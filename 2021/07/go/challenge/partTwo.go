package challenge

import (
	"math"
)

func PartTwo(instr string) int {
	inputSlice := parse(instr)
	meanTotal := 0
	for _, i := range inputSlice {
		meanTotal += i
	}
	floor := int(math.Floor(float64(meanTotal) / float64(len(inputSlice))))
	ceil := floor + 1
	fTotal, cTotal := 0, 0
	for _, v := range inputSlice {
		// calc floor mean total
		abs := floor - v
		if abs < 0 {
			abs = -abs
		}
		fTotal += (abs * (abs + 1)) / 2
		// calc ceil mean total
		abs = ceil - v
		if abs < 0 {
			abs = -abs
		}
		cTotal += (abs * (abs + 1)) / 2
	}
	// return min
	if cTotal < fTotal {
		return cTotal
	}
	return fTotal
}
