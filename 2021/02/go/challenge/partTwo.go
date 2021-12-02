package challenge

func PartTwo(instr string) int {
	instructions := parse(instr)
	depth, forward, aim := 0, 0, 0
	for _, instruction := range instructions {
		switch instruction.Direction {
		case "forward":
			forward += instruction.Value
			depth += aim * instruction.Value
		case "up":
			// depth -= instruction.Value
			aim -= instruction.Value
		case "down":
			// depth += instruction.Value
			aim += instruction.Value
		}
	}
	return depth * forward
}
