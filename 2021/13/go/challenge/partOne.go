package challenge

func PartOne(instr string) int {
	paper := parse(instr)
	if paper.Folds[0].Up {
		paper.foldUp(paper.Folds[0].Axis)
	} else {
		paper.foldLeft(paper.Folds[0].Axis)
	}
	total := 0
	for _, vy := range paper.Dots {
		total += len(vy)
	}
	return total
}
