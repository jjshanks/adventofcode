package challenge

func PartOne(instr string) int {
	inputSlice := parse(instr)
	sum := 0
	for _, b := range inputSlice {
		sum += b.PaperNeeded()
	}
	return sum
}
