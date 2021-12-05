package challenge

import (
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Vent struct {
	To   Point
	From Point
}

func parse(instr string) []Vent {
	r := regexp.MustCompile(`([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)`)
	lines := strings.Split(instr, "\n")
	vents := make([]Vent, 0, len(lines))
	for _, line := range lines {
		parts := r.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(parts[1])
		y1, _ := strconv.Atoi(parts[2])
		x2, _ := strconv.Atoi(parts[3])
		y2, _ := strconv.Atoi(parts[4])
		vents = append(vents, Vent{
			To: Point{
				X: x1,
				Y: y1,
			},
			From: Point{
				X: x2,
				Y: y2,
			},
		})
	}
	return vents
}

func fillStriaght(vent Vent, ventMap map[int]map[int]int) (total int) {
	minX, maxX, minY, maxY := vent.To.X, vent.From.X, vent.To.Y, vent.From.Y
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	if minY > maxY {
		minY, maxY = maxY, minY
	}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if _, ok := ventMap[x]; !ok {
				ventMap[x] = make(map[int]int)
			}
			ventMap[x][y] += 1
			if ventMap[x][y] == 2 {
				total += 1
			}
		}
	}
	return
}

func fillDiagonal(vent Vent, ventMap map[int]map[int]int) (total int) {
	xDelta, yDelta := 0, 0
	if vent.From.X > vent.To.X {
		xDelta = -1
	} else {
		xDelta = 1
	}
	if vent.From.Y > vent.To.Y {
		yDelta = -1
	} else {
		yDelta = 1
	}
	for x, y := vent.From.X, vent.From.Y; x != vent.To.X && y != vent.To.Y; {
		if _, ok := ventMap[x]; !ok {
			ventMap[x] = make(map[int]int)
		}
		ventMap[x][y] += 1
		if ventMap[x][y] == 2 {
			total += 1
		}
		x += xDelta
		y += yDelta
	}
	// update To coords
	if _, ok := ventMap[vent.To.X]; !ok {
		ventMap[vent.To.X] = make(map[int]int)
	}
	ventMap[vent.To.X][vent.To.Y] += 1
	if ventMap[vent.To.X][vent.To.Y] == 2 {
		total += 1
	}
	return
}
