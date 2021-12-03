package challenge

func PartTwo(instr string) int {
	values := parse(instr)
	oxygen := make([][]bool, len(values))
	co2 := make([][]bool, len(values))
	copy(oxygen, values)
	copy(co2, values)

	ox := findOxygen(oxygen)
	co := findCo2(co2)
	return co * ox
}

func findOxygen(oxygen [][]bool) int {
	return findValue(oxygen, true)
}

func findCo2(co2 [][]bool) int {
	return findValue(co2, false)
}

func findValue(in [][]bool, getMax bool) int {
	idx := 0
	for len(in) > 1 {
		ones := make([][]bool, 0)
		zeros := make([][]bool, 0)
		for _, o := range in {
			if o[idx] {
				ones = append(ones, o)
			} else {
				zeros = append(zeros, o)
			}
		}
		if getMax {
			if len(ones) >= len(zeros) {
				in = ones
			} else {
				in = zeros
			}
		} else {
			if len(zeros) <= len(ones) {
				in = zeros
			} else {
				in = ones
			}
		}
		idx += 1
	}
	return convert(in[0])
}
