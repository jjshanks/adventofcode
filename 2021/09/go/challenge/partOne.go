package challenge

func PartOne(instr string) int {
	heights := parse(instr)
	risk := 0
	for x := 0; x < len(heights); x += 1 {
		for y := 0; y < len(heights[x]); y += 1 {
			if isLowPoint(heights, x, y) {
				risk += 1 + int(heights[x][y])
			}
		}
	}
	return risk
}

func isLowPoint(heigths [][]byte, x, y int) bool {
	maxX := len(heigths)
	maxY := len(heigths[0])
	checking := heigths[x][y]
	if x+1 < maxX && heigths[x+1][y] <= checking {
		return false
	}
	if x-1 >= 0 && heigths[x-1][y] <= checking {
		return false
	}
	if y+1 < maxY && heigths[x][y+1] <= checking {
		return false
	}
	if y-1 >= 0 && heigths[x][y-1] <= checking {
		return false
	}
	return true
}
