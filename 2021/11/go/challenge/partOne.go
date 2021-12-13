package challenge

func PartOne(instr string) int {
	m := parse(instr)
	total := 0
	for i := 0; i < 100; i += 1 {
		total += step(m)
	}
	return total
}

type Point struct {
	X int
	Y int
}
