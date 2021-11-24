package challenge

import "strings"

func PartOne(instr string) int {
	total := 0
	for _, s := range strings.Split(instr, "\n") {
		if IsNice(s) {
			total += 1
		}
	}
	return total
}
