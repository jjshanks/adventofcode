package challenge

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Draws  []int
	Boards []*Board
}

type Board struct {
	Spaces   [][]int
	Marks    [][]bool
	RowMarks []int
	ColMarks []int
}

func (b *Board) Mark(num int) {
	for i, row := range b.Spaces {
		for j, space := range row {
			if space == num {
				b.Marks[i][j] = true
				b.RowMarks[i] += 1
				b.ColMarks[j] += 1
				return
			}
		}
	}
}

func (b *Board) IsWinner() bool {
	for i := 0; i < 5; i += 1 {
		if b.RowMarks[i] == 5 || b.ColMarks[i] == 5 {
			return true
		}
	}
	return false
}

func (b *Board) Score() int {
	total := 0
	for i := 0; i < 5; i += 1 {
		for j := 0; j < 5; j += 1 {
			if !b.Marks[i][j] {
				total += b.Spaces[i][j]
			}
		}
	}
	return total
}

func (b *Board) String() string {
	return fmt.Sprintf("%v ++ %v", b.Spaces, b.Marks)
}

func parse(instr string) *Game {
	lines := strings.Split(instr, "\n")
	drawStr := strings.Split(lines[0], ",")
	draw := make([]int, len(drawStr))
	for x, s := range drawStr {
		i, _ := strconv.Atoi(s)
		draw[x] = i
	}
	boards := make([]*Board, 0)
	for i := 0; i < len(lines)/6; i += 1 {
		boards = append(boards, &Board{})
		boards[i].Spaces = make([][]int, 5)
		boards[i].Marks = make([][]bool, 5)
		boards[i].ColMarks = make([]int, 5)
		boards[i].RowMarks = make([]int, 5)
		for j := 0; j < 5; j += 1 {
			boards[i].Spaces[j] = make([]int, 5)
			boards[i].Marks[j] = make([]bool, 5)
			lineIdx := i*6 + j + 2
			for k := 0; k < 5; k += 1 {
				s := lines[lineIdx]
				n, _ := strconv.Atoi(strings.Trim(s[(k*3):(k*3+2)], " "))
				boards[i].Spaces[j][k] = n
			}
		}
	}
	return &Game{
		Draws:  draw,
		Boards: boards,
	}
}
