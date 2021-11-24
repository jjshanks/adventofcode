package challenge

func PartTwo(instr string) int {
	instructions := parse(instr)
	grid := NewGrid()
	var fn func(g [][]int, x, y int)
	for _, i := range instructions {
		switch i.Code {
		case Toggle:
			fn = func(g [][]int, x, y int) {
				g[x][y] += 2
			}
		case TurnOn:
			fn = func(g [][]int, x, y int) {
				g[x][y] += 1
			}
		case Turnoff:
			fn = func(g [][]int, x, y int) {
				g[x][y] = max(0, g[x][y]-1)
			}
		}
		for x := i.Start.X; x <= i.End.X; x++ {
			for y := i.Start.Y; y <= i.End.Y; y++ {
				fn(grid, x, y)
			}
		}
	}
	sum := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			sum += grid[x][y]
		}
	}
	return sum
}
