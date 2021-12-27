package challenge

import (
	"strings"
)

type TwoDMap struct {
	Values  []byte
	MinCost []int
	Side    int
}

type Path struct {
	X    int
	Y    int
	D    int
	Past []Path
}

type Paths []Path

func (p Paths) Len() int           { return len(p) }
func (p Paths) Less(i, j int) bool { return p[i].D < p[j].D }
func (p Paths) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *Paths) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*p = append(*p, x.(Path))
}

func (p *Paths) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (m *TwoDMap) HasPoint(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= m.Side || y >= m.Side {
		return false
	}
	return true
}

func (m *TwoDMap) At(x, y int) byte {
	return m.Values[x*m.Side+y]
}

func (m *TwoDMap) Set(x, y int, v byte) {
	if m.HasPoint(x, y) {
		m.Values[x*m.Side+y] = v
	}
}

func (m *TwoDMap) IsEnd(x, y int) bool {
	return x == m.Side-1 && y == m.Side-1
}

func (m *TwoDMap) IsLowest(x, y, v int) bool {
	return m.MinCost[x*m.Side+y] == 0 || m.MinCost[x*m.Side+y] > v
}

func (m *TwoDMap) SetLowest(x, y, v int) {
	m.MinCost[x*m.Side+y] = v
}

func (m *TwoDMap) Neighbors(p Path) []Path {
	cost := p.D + int(m.At(p.X, p.Y))
	if !m.IsLowest(p.X, p.Y, cost) {
		return []Path{}
	}
	m.SetLowest(p.X, p.Y, cost)
	paths := make([]Path, 0, 4)
	if m.HasPoint(p.X-1, p.Y) {
		paths = append(paths, Path{X: p.X - 1, Y: p.Y, D: cost})
	}
	if m.HasPoint(p.X+1, p.Y) {
		paths = append(paths, Path{X: p.X + 1, Y: p.Y, D: cost})
	}
	if m.HasPoint(p.X, p.Y+1) {
		paths = append(paths, Path{X: p.X, Y: p.Y + 1, D: cost})
	}
	if m.HasPoint(p.X, p.Y-1) {
		paths = append(paths, Path{X: p.X, Y: p.Y - 1, D: cost})
	}
	return paths
}

func expand(m *TwoDMap) *TwoDMap {
	newSide := m.Side * 5
	values := make([]byte, newSide*newSide)
	minCost := make([]int, newSide*newSide)
	newMap := &TwoDMap{
		Values:  values,
		MinCost: minCost,
		Side:    newSide,
	}
	for x := 0; x < m.Side; x += 1 {
		for y := 0; y < m.Side; y += 1 {
			for i := byte(0); i < 5; i += 1 {
				for j := byte(0); j < 5; j += 1 {
					v := m.At(x, y) + i + j
					for v >= 10 {
						v -= 9
					}
					newMap.Set(m.Side*int(i)+x, m.Side*int(j)+y, v)
				}
			}
		}
	}
	return newMap
}

func parse(instr string) *TwoDMap {
	lines := strings.Split(instr, "\n")
	values := make([]byte, 0, len(lines)*len(lines))
	minCost := make([]int, len(lines)*len(lines))

	for _, line := range lines {
		for _, b := range line {
			values = append(values, byte(b)-48)
		}
	}
	return &TwoDMap{
		Values:  values,
		MinCost: minCost,
		Side:    len(lines),
	}
}
