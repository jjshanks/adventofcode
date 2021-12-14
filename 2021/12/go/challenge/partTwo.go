package challenge

import (
	"strings"
)

func PartTwo(instr string) int {
	cave := parse(instr)
	visited := make([]bool, len(cave.V))
	paths := 0
	visited[cave.Start] = true
	path := &Path{
		Visited: clone(visited),
		Vertex:  cave.Start,
		N:       []string{"start"},
	}
	queue := []*Path{path}
	seen := make(map[string]bool)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		cave.G.Visit(current.Vertex, func(w int, _ int64) bool {
			isSmall := strings.ToLower(cave.V[w]) == cave.V[w]
			if cave.End == w {
				path := strings.Join(current.N, ",")
				if _, ok := seen[path]; ok {
					return false
				}
				seen[path] = true
				paths += 1
				return false
			}
			if !current.Visited[w] {
				newVisit := clone(current.Visited)
				newNode := cloneS(current.N)
				newVisit[w] = true && isSmall
				path := &Path{
					Visited: newVisit,
					Vertex:  w,
					D:       current.D,
					N:       append(newNode, cave.V[w]),
				}
				queue = append(queue, path)
			}
			if isSmall && !current.D && !current.Visited[w] {
				newVisit := clone(current.Visited)
				newNode := cloneS(current.N)

				path := &Path{
					Visited: newVisit,
					Vertex:  w,
					D:       true,
					N:       append(newNode, cave.V[w]),
				}
				queue = append(queue, path)
			}
			return false
		})
	}
	return paths
}
