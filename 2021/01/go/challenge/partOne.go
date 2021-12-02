package challenge

func PartOne(instr string) int {
	values := parse(instr)
	inc := 0
	for i := 0; i < len(values)-1; i++ {
		if values[i] < values[i+1] {
			inc += 1
		}
	}
	return inc
}
