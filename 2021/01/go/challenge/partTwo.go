package challenge

func PartTwo(instr string) int {
	values := parse(instr)
	inc := 0
	for i := 0; i < len(values)-3; i++ {
		a := values[i] + values[i+1] + values[i+2]
		b := a - values[i] + values[i+3]
		if a < b {
			inc += 1
		}
	}
	return inc
}
