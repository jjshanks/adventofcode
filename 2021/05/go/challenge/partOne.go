package challenge

import (
	"fmt"
)

func PartOne(instr string) int {
	vents := parse(instr)
	fmt.Println()
	ventMap := make(map[int]map[int]int)
	total := 0
	for _, vent := range vents {
		if vent.To.X != vent.From.X && vent.To.Y != vent.From.Y {
			continue
		}
		total += fillStriaght(vent, ventMap)
	}
	return total
}
