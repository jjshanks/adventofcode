package challenge

import "strings"

func PartOne(instr string) int {
	cave := parse(instr)
	visited := make([]bool, len(cave.V))
	paths := 0
	visited[cave.Start] = true
	path := &Path{
		Visited: clone(visited),
		Vertex:  cave.Start,
	}
	queue := []*Path{path}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		cave.G.Visit(current.Vertex, func(w int, _ int64) bool {
			if cave.End == w {
				paths += 1
				return false
			}
			if !current.Visited[w] {
				newVisit := clone(current.Visited)
				isSmall := strings.ToLower(cave.V[w]) == cave.V[w]
				newVisit[w] = true && isSmall
				path := &Path{
					Visited: newVisit,
					Vertex:  w,
				}
				queue = append(queue, path)
			}
			return false
		})
	}
	return paths
}
