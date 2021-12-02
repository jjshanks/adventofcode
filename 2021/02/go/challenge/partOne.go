package challenge

func PartOne(instr string) int {
	instructions := parse(instr)
	depth, forward := 0, 0
	for _, instruction := range instructions {
		switch instruction.Direction {
		case "forward":
			forward += instruction.Value
		case "up":
			depth -= instruction.Value
		case "down":
			depth += instruction.Value
		}
	}
	return depth * forward
}
