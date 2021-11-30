package challenge

func PartTwo(instr string) int {
	result := instr
	for i := 0; i < 50; i++ {
		result = convert(result)
	}
	return len(result)
}
