package challenge

func PartOne(instr string) int {
	floor := 0
	for _, c := range instr {
		if c == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return floor
}
