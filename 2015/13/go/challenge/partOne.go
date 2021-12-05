package challenge

func PartOne(instr string) int {
	seatings := parse(instr)
	names := make([]string, 0, len(seatings))
	for k := range seatings {
		names = append(names, k)
	}

	return findMaxHappiness(seatings, names)
}
