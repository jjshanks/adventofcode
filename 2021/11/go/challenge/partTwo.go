package challenge

func PartTwo(instr string) int {
	total := 0
	steps := 0
	m := parse(instr)
	for total != 100 {
		total = step(m)
		steps += 1
	}
	return steps
}
