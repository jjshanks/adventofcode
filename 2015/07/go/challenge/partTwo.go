package challenge

import "fmt"

func PartTwo(instr string) int {
	instructions := parse(instr)
	wireValues := make(map[string]uint16)
	wireValues["b"] = 956
	fmt.Println()
	for {
		for _, instruction := range instructions {
			if _, done := wireValues[instruction.Out]; done {
				continue
			}
			out, err := instruction.eval(wireValues)
			if err != nil {
				continue
			}
			wireValues[instruction.Out] = out
		}
		val, done := wireValues["a"]
		if done {
			return int(val)
		}
	}
}
