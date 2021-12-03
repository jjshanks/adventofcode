package challenge

func PartOne(instr string) int {
	values := parse(instr)
	ones := make([]int, len(values[0]))
	zeros := make([]int, len(values[0]))
	for _, v := range values {
		for idx, b := range v {
			if b {
				ones[idx] += 1
			} else {
				zeros[idx] += 1
			}
		}
	}
	gamma := 0
	epsilon := 0
	for idx, _ := range ones {
		shift := len(ones) - idx - 1
		if ones[idx] > zeros[idx] {
			gamma += 1 << shift
		} else {
			epsilon += 1 << shift
		}
	}
	return gamma * epsilon
}
