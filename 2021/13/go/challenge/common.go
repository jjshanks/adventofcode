package challenge

import (
	"strconv"
	"strings"
)

type Paper struct {
	Dots  map[int]map[int]bool
	Folds []*Fold
	X     int
	Y     int
}

func (p *Paper) foldUp(axis int) {
	dots := make(map[int]map[int]bool)
	for y := range p.Dots {
		newY := axis - (y - axis)
		if y < axis {
			newY = y
		}
		for x := range p.Dots[y] {
			if _, ok := dots[newY]; !ok {
				dots[newY] = make(map[int]bool)
			}
			dots[newY][x] = p.Dots[y][x] || p.Dots[newY][x]
		}
	}
	p.Dots = dots
	p.Y = axis
}

func (p *Paper) foldLeft(axis int) {
	dots := make(map[int]map[int]bool)
	for y := range p.Dots {
		for x := range p.Dots[y] {
			newX := axis - (x - axis)
			if x < axis {
				newX = x
			}
			if _, ok := dots[y]; !ok {
				dots[y] = make(map[int]bool)
			}
			dots[y][newX] = p.Dots[y][x] || p.Dots[y][newX]
		}
	}
	p.Dots = dots
	p.X = axis
}

type Fold struct {
	Up   bool
	Axis int
}

func parse(instr string) *Paper {
	maxX, maxY := 0, 0
	dots := make(map[int]map[int]bool)
	lines := strings.Split(instr, "\n")
	idx := 0
	for ; idx < len(lines); idx += 1 {
		line := lines[idx]
		if line == "" {
			break
		}
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		if x > maxX {
			maxX = x
		}
		y, _ := strconv.Atoi(coords[1])
		if y > maxY {
			maxY = y
		}
		if _, ok := dots[y]; !ok {
			dots[y] = make(map[int]bool)
		}
		dots[y][x] = true
	}
	idx += 1
	folds := make([]*Fold, len(lines)-idx)
	for i := 0; i < len(lines)-idx; i += 1 {
		line := lines[idx+i]
		tokens := strings.Split(line, " ")
		parts := strings.Split(tokens[2], "=")
		axis, _ := strconv.Atoi(parts[1])
		up := true
		if parts[0] != "y" {
			up = false
		}
		folds[i] = &Fold{Up: up, Axis: axis}
	}
	return &Paper{
		Folds: folds,
		Dots:  dots,
		X:     maxX,
		Y:     maxY,
	}
}
