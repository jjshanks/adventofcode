package challenge

import (
	"strconv"
	"strings"
)

func PartTwo(instr string) int {
	total := 0
	lines := strings.Split(instr, "\n")
	for _, line := range lines {
		minRed := int64(0)
		minBlu := int64(0)
		minGrn := int64(0)
		draws := strings.Split(strings.Split(line, ":")[1], ";")
		for _, draw := range draws {
			cubes := strings.Split(draw, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				parts := strings.Split(cube, " ")
				color := parts[1]
				num, _ := strconv.ParseInt(parts[0], 10, 32)
				switch {
				case color == "red":
					if num > minRed {
						minRed = num
					}
				case color == "blue":
					if num > minBlu {
						minBlu = num
					}
				case color == "green":
					if num > minGrn {
						minGrn = num
					}
				}
			}
		}
		total = total + (int(minBlu) * int(minGrn) * int(minRed))
	}
	return total
}
