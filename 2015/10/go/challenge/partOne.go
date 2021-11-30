package challenge

func PartOne(instr string) int {
	result := instr
	for i := 0; i < 40; i++ {
		result = convert(result)
	}
	return len(result)
}
