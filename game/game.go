package game

import (
	"github.com/deadly990/gominesweeper/generation"
)

type Game struct {
	board    generation.Board
	Revealed [][]int
}

func NewGame(board generation.Board) *Game {
	width, height := generation.BoardSize(board)

	var revealed = make([][]int, height)
	for i := 0; i < height; i++ {
		revealed[i] = make([]int, width)
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

func tileValue(game Game, coord Coordinate) *int {
	return &(game.Revealed[coord.y][coord.x])
}

func Move(game Game, y int, x int) {
	var reveal = func(y int, x int) {
		switch value := game.Revealed[y][x]; value {
		case -10:
			*tileValue(game, Coordinate{x, y}) = 0
		case -1, -2, -3, -4, -5, -6, -7, -8, -9:
			*tileValue(game, Coordinate{x, y}) = -value
		}
	}
	var addNeighbors = func(list *[]Coordinate, origin Coordinate) {
		for _, coord := range Adjacent(origin) {
			if generation.IsInRange(game.board, coord.y, coord.x) && *tileValue(game, coord) < 0 {
				*list = append(*list, coord)
			}
		}
	}
	var queue []Coordinate
	queue = append(queue, Coordinate{x, y})
	for len(queue) > 0 {
		queuedCoord := queue[0]
		queue = queue[1:]
		if *tileValue(game, queuedCoord) == -10 {
			addNeighbors(&queue, queuedCoord)
		}
		reveal(queuedCoord.y, queuedCoord.x)
	}
}
