package challenge

func PartOne(instr string) int {
	game := parse(instr)
	for _, draw := range game.Draws {
		for _, board := range game.Boards {
			board.Mark(draw)
			if board.IsWinner() {
				return board.Score() * draw
			}
		}
	}
	return 0
}
