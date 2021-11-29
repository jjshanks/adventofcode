package challenge

func PartOne(instr string) int {
	vertexes := parse(instr)
	min := int(^uint(0) >> 1)
	visited := make(map[string]bool)
	for _, v := range vertexes {
		visited[v.Name] = false
	}
	for _, start := range vertexes {
		result := eval(start, visited, int(^uint(0)>>1), func(i, j int) bool { return i > j })
		if result < min {
			min = result
		}
	}
	return min
}
