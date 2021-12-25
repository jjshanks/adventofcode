package challenge

func PartTwo(instr string) int {
	pairs, template := parse(instr)
	return process(template, pairs, 40)
}
