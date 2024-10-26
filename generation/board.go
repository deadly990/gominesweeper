package generation

import (
	"math/rand"
)

type Board struct {
	Mines int
	Field [][]int
	seed  int64
}

func NewBoard(mines int, width int, height int, seed int64) *Board {
	board := Board{mines, blankField(width, height), seed}
	generateMines(board)
	return &board
}

func blankField(width int, height int) [][]int {
	var arr = make([][]int, height)
	for i := 0; i < width; i++ {
		arr[i] = make([]int, width)
	}
	return arr
}

func isInRange(board Board, x int, y int) bool {
	var width = len(board.Field)
	var height = len(board.Field[0])
	return x >= 0 && x < width && y >= 0 && y < height
}

// Returns false if the tile is out of bounds or is a mine.
func isValidTile(board Board, x int, y int) bool {
	return isInRange(board, x, y) && board.Field[x][y] != -1
}

func generateMines(board Board) {
	var width = len(board.Field)
	var height = len(board.Field[0])

	// Traverses all adjacent tiles in a 1 tile radius
	var populateHints = func(x int, y int) {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			for yOffset := -1; yOffset <= 1; yOffset++ {
				if isValidTile(board, x+xOffset, y+yOffset) {
					board.Field[x+xOffset][y+yOffset] += 1
				}
			}
		}
	}

	var random = rand.New(rand.NewSource((board.seed)))
	// Iterates until n mines have been successfully placed.
	for count := 0; count < board.Mines; count++ {
		var x = random.Intn(width)  // X offset
		var y = random.Intn(height) // Y offset
		if board.Field[x][y] == -1 {
			count--
			continue
			// Does not count to the progress of mines on the occasion that a mine already exists in a location.
		}
		board.Field[x][y] = -1
		populateHints(x, y)
	}
}

func Validate(board Board) bool {
	var width = len(board.Field)
	var height = len(board.Field[0])

	var mineCount = func() bool {
		var actual = 0
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if board.Field[x][y] == -1 {
					actual++
				}
			}
		}
		return actual == board.Mines
	}

	var hintVeracity = func() bool {
		var countSurroundings = func(x int, y int) int {

			var minesFound = 0
			if board.Field[x][y] == -1 {
				return -1
			}
			for xOffset := -1; xOffset <= 1; xOffset++ {
				for yOffset := -1; yOffset <= 1; yOffset++ {
					if isInRange(board, x+xOffset, y+yOffset) {
						if board.Field[x+xOffset][y+yOffset] == -1 {
							minesFound++
						}
					}
				}
			}
			return minesFound
		}

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if board.Field[x][y] != countSurroundings(x, y) {
					return false
				}
			}
		}
		return true
	}

	return mineCount() && hintVeracity()
}
