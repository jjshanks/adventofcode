package challenge

import "fmt"

func PartOne(instr string) int {
	x, y := 0, 0
	visited := make(map[string]bool)
	visited["0--0"] = true
	for _, d := range instr {
		switch d {
		case 'v':
			y -= 1
		case '^':
			y += 1
		case '<':
			x -= 1
		case '>':
			x += 1
		}
		key := fmt.Sprintf("%d--%d", x, y)
		visited[key] = true
	}
	return len(visited)
}
