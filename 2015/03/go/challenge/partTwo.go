package challenge

import "fmt"

func PartTwo(instr string) int {
	sx, sy, rx, ry := 0, 0, 0, 0
	visited := make(map[string]bool)
	visited["0--0"] = true
	for i, d := range instr {
		switch d {
		case 'v':
			if i%2 == 0 {
				sy -= 1
			} else {
				ry -= 1
			}
		case '^':
			if i%2 == 0 {
				sy += 1
			} else {
				ry += 1
			}
		case '<':
			if i%2 == 0 {
				sx -= 1
			} else {
				rx -= 1
			}
		case '>':
			if i%2 == 0 {
				sx += 1
			} else {
				rx += 1
			}
		}
		santaKey := fmt.Sprintf("%d--%d", sx, sy)
		visited[santaKey] = true
		robotKey := fmt.Sprintf("%d--%d", rx, ry)
		visited[robotKey] = true
	}
	return len(visited)
}
