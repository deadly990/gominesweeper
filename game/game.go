package game

import (
	"github.com/deadly990/gominesweeper/generation"
)

type Game struct {
	Board    generation.Board
	Revealed [][]int
	Moves    []Coordinate
}

func NewGame(board generation.Board) *Game {
	width, height := board.BoardSize()

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

	return &Game{board, revealed, []Coordinate{}}
}

func (game *Game) tileValue(coord Coordinate) *int {
	return &(game.Revealed[coord.Y][coord.X])
}

func (game *Game) Move(coord Coordinate, action func(Coordinate)) {
	action(coord)
	game.Moves = append(game.Moves, coord)
}

// Clears a tile at position (x, y)
func (game *Game) Clear(coord Coordinate) {
	var queue []Coordinate
	queue = append(queue, coord)
	for len(queue) > 0 {
		queuedCoord := queue[0]
		queue = queue[1:]
		if *game.tileValue(queuedCoord) == -10 {
			for _, adjacent := range queuedCoord.Adjacent() {
				if game.isValidClear(adjacent) {
					queue = append(queue, adjacent)
				}
			}
		}
		game.revealTileValue(queuedCoord)
	}
}

func (game *Game) isValidClear(coord Coordinate) bool {
	return game.Board.IsInRange(coord.Y, coord.X) && !game.isRevealed(coord)
}

func (game *Game) isRevealed(coord Coordinate) bool {
	return *game.tileValue(coord) >= 0
}

func (game *Game) revealTileValue(coord Coordinate) {
	switch value := game.Revealed[coord.Y][coord.X]; value {
	case -10: // Unrevealed blank tile.
		*game.tileValue(coord) = 0
	case -1, -2, -3, -4, -5, -6, -7, -8, -9:
		*game.tileValue(coord) = -value
	}
}
