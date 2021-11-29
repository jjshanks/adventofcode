package challenge

func PartTwo(instr string) int {
	vertexes := parse(instr)
	max := 0
	visited := make(map[string]bool)
	for _, v := range vertexes {
		visited[v.Name] = false
	}
	for _, start := range vertexes {
		result := eval(start, visited, 0, func(i, j int) bool { return i < j })
		if result > max {
			max = result
		}
	}
	return max
}
