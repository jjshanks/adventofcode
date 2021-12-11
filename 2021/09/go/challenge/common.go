package challenge

import "strings"

func parse(instr string) [][]byte {
	lines := strings.Split(instr, "\n")
	heights := make([][]byte, len(lines))
	for x, line := range lines {
		heights[x] = make([]byte, len(line))
		for y, c := range line {
			i := byte(c) - byte(48)
			heights[x][y] = i

		}
	}
	return heights
}
