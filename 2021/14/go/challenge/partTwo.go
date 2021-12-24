package challenge

func PartTwo(instr string) int {
	pairs, template := parse(instr)
	return newProcess(template, pairs, 40)
}
