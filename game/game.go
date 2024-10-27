package game

import (
	"github.com/deadly990/gominesweeper/generation"
)

type Game struct {
	Board    generation.Board
	Revealed [][]int
}

type coordinate struct {
	x int
	y int
}

func NewGame(board generation.Board) *Game {
	height, width := generation.BoardSize(board)

	var revealed = make([][]int, width)
	for i := 0; i < height; i++ {
		revealed[i] = make([]int, height)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch value := board.Field[y][x]; value {
			case 0:
				revealed[y][x] = -10
			case 1, 2, 3, 4, 5, 6, 7, 8:
				revealed[y][x] = -value
			case -9:
				revealed[y][x] = value
			}
		}
	}
	return &Game{board, revealed}
}
