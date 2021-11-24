package challenge

import "strings"

func PartTwo(instr string) int {
	total := 0
	for _, s := range strings.Split(instr, "\n") {
		if IsNicePartTwo(s) {
			total += 1
		}
	}
	return total
}
