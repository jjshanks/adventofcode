package challenge

type Point struct {
	X int
	Y int
}

func PartTwo(instr string) int {
	heights := parse(instr)
	visited := make([][]bool, len(heights))
	for x := 0; x < len(heights); x += 1 {
		visited[x] = make([]bool, len(heights[x]))
		for y := 0; y < len(heights[x]); y += 1 {
			if heights[x][y] == 9 {
				visited[x][y] = true
			}
		}
	}
	areas := make([]int, 0)
	for x := 0; x < len(heights); x += 1 {
		for y := 0; y < len(heights[x]); y += 1 {
			if visited[x][y] {
				continue
			}
			areas = append(areas, fill(heights, visited, x, y))
		}
	}
	first, areas := popMax(areas)
	second, areas := popMax(areas)
	third, _ := popMax(areas)
	return first * second * third
}

func popMax(areas []int) (int, []int) {
	max := 0
	var idx int
	for i, v := range areas {
		if v > max {
			max = v
			idx = i
		}
	}
	areas[idx] = areas[len(areas)-1]
	return max, areas[:len(areas)-1]
}

func fill(heights [][]byte, visited [][]bool, x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	if x >= len(heights) || y >= len(heights[x]) {
		return 0
	}
	if visited[x][y] {
		return 0
	}
	visited[x][y] = true
	return 1 +
		fill(heights, visited, x+1, y) +
		fill(heights, visited, x-1, y) +
		fill(heights, visited, x, y+1) +
		fill(heights, visited, x, y-1)
}
