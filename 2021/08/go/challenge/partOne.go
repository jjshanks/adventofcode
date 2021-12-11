package challenge

func PartOne(instr string) int {
	inputSlice := parse(instr)
	total := 0
	for _, o := range inputSlice {
		for _, d := range o.Digits {
			if len(d) <= 4 || len(d) == 7 {
				total += 1
			}
		}
	}
	return total
}
