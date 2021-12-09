package challenge

func PartOne(instr string) int {
	inputSlice := parse(instr)
	newBorn1, newBorn2 := 0, 0
	for i := 0; i < 80; i++ {
		dueIdx := i % 7
		matured := newBorn1
		newBorn1 = newBorn2
		newBorn2 = inputSlice[dueIdx]
		inputSlice[dueIdx] += matured
	}
	total := newBorn1 + newBorn2
	for _, v := range inputSlice {
		total += v
	}
	return total
}
