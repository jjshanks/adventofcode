package challenge

func PartTwo(instr string) int {
	floor := 0
	for n, c := range instr {
		if c == '(' {
			floor += 1
		} else {
			floor -= 1
		}
		if floor < 0 {
			return n + 1
		}
	}
	return 0
}
