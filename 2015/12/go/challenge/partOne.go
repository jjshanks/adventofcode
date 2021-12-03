package challenge

func PartOne(instr string) int {
	json := parse(instr)
	total := processNode(json, "")
	return total
}
