package challenge

func PartTwo(instr string) int {
	game := parse(instr)
	for _, draw := range game.Draws {
		winners := make([]*Board, 0)
		for _, board := range game.Boards {
			board.Mark(draw)
			if board.IsWinner() && len(game.Boards) == 1 {
				return board.Score() * draw
			}
			if board.IsWinner() {
				winners = append(winners, board)
			}
		}
		for _, r := range winners {
			for i, v := range game.Boards {
				if r == v {
					game.Boards = append(game.Boards[:i], game.Boards[i+1:]...)
				}
			}
		}
	}
	return 0
}
