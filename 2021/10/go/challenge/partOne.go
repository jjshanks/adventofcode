package challenge

func PartOne(instr string) int {
	total := 0
	lines := parse(instr)
	for _, line := range lines {
		q := make([]rune, 0, len(line))
		for _, c := range line {
			if contains(closers, c) {
				opener := q[len(q)-1]
				if closerPairs[c] == opener {
					q = q[:len(q)-1]
					continue
				}
				total += closerValuesPart1[c]
				break
			}
			q = append(q, c)
		}
	}
	return total
}
