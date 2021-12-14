package challenge

import (
	"strings"

	"github.com/yourbasic/graph"
)

type Caves struct {
	G     *graph.Mutable
	V     []string
	Start int
	End   int
}

type Path struct {
	Visited []bool
	Vertex  int
	D       bool
	N       []string
}

func clone(s []bool) []bool {
	d := make([]bool, len(s))
	copy(d, s)
	return d
}

func cloneS(s []string) []string {
	d := make([]string, len(s))
	copy(d, s)
	return d
}

func parse(instr string) *Caves {
	vertMap := make(map[string]int)
	vertList := make([]string, 0)
	for _, line := range strings.Split(instr, "\n") {
		verts := strings.Split(line, "-")
		if _, ok := vertMap[verts[0]]; !ok {
			vertList = append(vertList, verts[0])
			vertMap[verts[0]] = len(vertMap)
		}
		if _, ok := vertMap[verts[1]]; !ok {
			vertList = append(vertList, verts[1])
			vertMap[verts[1]] = len(vertMap)
		}
	}
	g := graph.New(len(vertList))
	for _, line := range strings.Split(instr, "\n") {
		verts := strings.Split(line, "-")
		v := vertMap[verts[0]]
		w := vertMap[verts[1]]
		g.AddBoth(v, w)
	}
	return &Caves{G: g, V: vertList, Start: vertMap["start"], End: vertMap["end"]}
}
