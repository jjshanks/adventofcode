package challenge

func PartTwo(instr string) int {
	json := parse(instr)
	total := processNode(json, "red")
	return total
}
