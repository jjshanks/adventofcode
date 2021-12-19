package challenge

import "fmt"

func PartTwo(instr string) int {
	paper := parse(instr)
	for _, fold := range paper.Folds {
		if fold.Up {
			paper.foldUp(fold.Axis)
		} else {
			paper.foldLeft(fold.Axis)
		}
	}
	fmt.Println()
	for y := 0; y < paper.Y; y += 1 {
		for x := 0; x < paper.X; x += 1 {
			if paper.Dots[y][x] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return 0
}
