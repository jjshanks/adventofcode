package challenge

import (
	"regexp"
)

func PartOne(instr string) int {
	lines := parse(instr)
	r := regexp.MustCompile("\\\\(\"|\\\\|x[0-9a-f]{2})")
	all, sub := 0, 0
	for _, line := range lines {
		all += len(line)
		stripped := line[1 : len(line)-1]
		simple := r.ReplaceAllString(stripped, "x")
		sub += len(simple)
	}
	return all - sub
}
