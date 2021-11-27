package challenge

import (
	"fmt"
)

func PartTwo(instr string) int {
	lines := parse(instr)
	all, extra := 0, 0
	for _, line := range lines {
		all += len(line)
		quoted := fmt.Sprintf("%q", line)
		extra += len(quoted)
	}
	return extra - all
}
