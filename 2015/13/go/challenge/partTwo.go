package challenge

func PartTwo(instr string) int {
	seatings := parse(instr)
	names := make([]string, 0, len(seatings))
	seatings["me"] = make(map[string]int)
	for k := range seatings {
		names = append(names, k)
		seatings["me"][k] = 0
		seatings[k]["me"] = 0
	}

	return findMaxHappiness(seatings, names)
}
