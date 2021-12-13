package challenge

import "strings"

func parse(instr string) [][]byte {
	lines := strings.Split(instr, "\n")
	m := make([][]byte, len(lines))
	for i, line := range lines {
		m[i] = make([]byte, len(line))
		for j, r := range line {
			m[i][j] = byte(r) - byte(48)
		}
	}
	return m
}

func step(m [][]byte) int {
	total := 0
	flashed := make([][]bool, len(m))
	toFlash := make([]*Point, 0)
	for i, r := range m {
		flashed[i] = make([]bool, len(r))
	}
	for x := 0; x < len(m); x += 1 {
		for y := 0; y < len(m[x]); y += 1 {
			m[x][y] += 1
			if m[x][y] > 9 {
				toFlash = append(toFlash, &Point{X: x, Y: y})
			}
		}
	}
	for len(toFlash) > 0 {
		p := toFlash[len(toFlash)-1]
		toFlash = toFlash[:len(toFlash)-1]
		if flashed[p.X][p.Y] {
			continue
		}
		flashed[p.X][p.Y] = true
		extra := flashNeighbors(m, flashed, p)
		toFlash = append(toFlash, extra...)
	}
	for x := 0; x < len(m); x += 1 {
		for y := 0; y < len(m[x]); y += 1 {
			if flashed[x][y] {
				m[x][y] = 0
				total += 1
			}
		}
	}
	return total
}

func flashNeighbors(m [][]byte, f [][]bool, p *Point) []*Point {
	toFlash := make([]*Point, 0)
	for x := -1; x <= 1; x += 1 {
		for y := -1; y <= 1; y += 1 {
			if p.X+x < 0 || p.Y+y < 0 {
				continue
			}
			if p.X+x >= len(m) || p.Y+y >= len(m[p.X+x]) {
				continue
			}
			m[p.X+x][p.Y+y] += 1
			if m[p.X+x][p.Y+y] > 9 && !f[p.X+x][p.Y+y] {
				toFlash = append(toFlash, &Point{X: p.X + x, Y: p.Y + y})
			}
		}
	}
	return toFlash
}
