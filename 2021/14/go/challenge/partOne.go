package challenge

func PartOne(instr string) int {
	pairs, template := parse(instr)
	return process(template, pairs, 10)
}
