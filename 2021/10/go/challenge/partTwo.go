package challenge

import "sort"

func PartTwo(instr string) int {
	lines := parse(instr)
	scores := make([]int, 0, len(lines))
	for _, line := range lines {
		q := make([]rune, 0, len(line))
		invalid := false
		for _, c := range line {
			if contains(closers, c) {
				opener := q[len(q)-1]
				if closerPairs[c] == opener {
					q = q[:len(q)-1]
					continue
				}
				invalid = true
				break
			}
			q = append(q, c)
		}
		if !invalid {
			score := 0
			for len(q) > 0 {
				opener := q[len(q)-1]
				q = q[:len(q)-1]
				closer := openerPairs[opener]
				score = score*5 + closerValuesPart2[closer]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	idx := len(scores) / 2
	return scores[idx]
}
