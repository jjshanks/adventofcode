package challenge

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(instr string) []*Vertex {
	vertexMap := make(map[string]*Vertex)
	// create vertexes
	for _, line := range strings.Split(instr, "\n") {
		parts := strings.Split(line, " ")
		if _, ok := vertexMap[parts[0]]; !ok {
			vertex := &Vertex{
				Name: parts[0],
			}
			vertexMap[parts[0]] = vertex
		}
		if _, ok := vertexMap[parts[2]]; !ok {
			vertex := &Vertex{
				Name: parts[2],
			}
			vertexMap[parts[2]] = vertex
		}
	}
	// create edges
	for _, line := range strings.Split(instr, "\n") {
		parts := strings.Split(line, " ")
		distance, _ := strconv.Atoi(parts[4])
		from := vertexMap[parts[0]]
		to := vertexMap[parts[2]]
		edge := &Edge{
			From:     from,
			To:       to,
			Distance: distance,
		}
		from.Outs = append(from.Outs, edge)
		to.Ins = append(to.Ins, edge)
	}
	vertexes := make([]*Vertex, 0, len(vertexMap))
	for _, v := range vertexMap {
		vertexes = append(vertexes, v)
	}
	return vertexes
}

type Edge struct {
	To       *Vertex
	From     *Vertex
	Distance int
}

func copy(org map[string]bool) map[string]bool {
	copy := make(map[string]bool)
	for k, v := range org {
		copy[k] = v
	}
	return copy
}

func (e *Edge) String() string {
	return fmt.Sprintf("%s to %s in %d", e.To.Name, e.From.Name, e.Distance)
}

type Candidate struct {
	Vertex   *Vertex
	Distance int
}

type Vertex struct {
	Name string
	Ins  []*Edge
	Outs []*Edge
}

func (v *Vertex) String() string {
	return fmt.Sprintf("Name: %s\n  ins: %v\n  outs: %v\n", v.Name, v.Ins, v.Outs)
}

func eval(current *Vertex, visited map[string]bool, start int, betterFn func(int, int) bool) int {
	dist, _ := search(current, 0, visited, start, betterFn)
	return dist
}

func search(current *Vertex, total int, visited map[string]bool, start int, betterFn func(int, int) bool) (int, error) {
	localVisited := copy(visited)
	localVisited[current.Name] = true
	visitedAll := true
	for _, v := range localVisited {
		visitedAll = visitedAll && v
	}
	if visitedAll {
		return total, nil
	}
	cur := start //
	candidates := make([]Candidate, 0)
	for _, vertex := range current.Ins {
		candidates = append(candidates, Candidate{
			Vertex:   vertex.From,
			Distance: vertex.Distance,
		})
	}
	for _, vertex := range current.Outs {
		candidates = append(candidates, Candidate{
			Vertex:   vertex.To,
			Distance: vertex.Distance,
		})
	}
	for _, next := range candidates {
		if visited[next.Vertex.Name] {
			continue
		}
		innerVisited := copy(localVisited)
		dist, err := search(next.Vertex, total+next.Distance, innerVisited, start, betterFn)
		if err == nil && betterFn(cur, dist) {
			cur = dist
		}
	}
	return cur, nil
}
