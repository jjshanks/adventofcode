package challenge

func PartTwo(instr string) int {
	inputSlice := parse(instr)
	sum := 0
	for _, b := range inputSlice {
		sum += b.RibbonNeeded()
	}
	return sum
}
