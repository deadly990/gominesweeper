package generation

import (
	"fmt"
	"math/rand"
)

type Board struct {
	Mines int
	Field [][]int
	seed  int64
}

func NewBoard(mines int, width int, height int, seed int64) (*Board, error) {
	var inputValidation = func() error {
		if mines < 0 {
			return fmt.Errorf("mines value cannot be negative")
		}
		if width < 1 || height < 1 {
			return fmt.Errorf("width and height must be greater than or equal to 1. Actual: %dx%d", width, height)
		}
		return nil
	}

	if inputErr := inputValidation(); inputErr != nil {
		return nil, inputErr
	}

	board := Board{mines, blankField(width, height), seed}
	var genErr = generateMines(board)
	return &board, genErr
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

func generateMines(board Board) error {
	var width = len(board.Field)
	var height = len(board.Field[0])

	var pregenTests = func() error {
		if board.Mines > width*height {
			return fmt.Errorf("MineOverload: The board size specified cannot hold the number mines provided")
		}
		if !Validate(board) {
			return fmt.Errorf("InvalidBoard: Board generation failed, created an invalid board")
		}
		return nil // No error detected
	}
	if testResult := pregenTests(); testResult != nil {
		return testResult
	}

	// Traverses all adjacent tiles in a 1 tile radius
	var populateHints = func(x int, y int) {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			for yOffset := -1; yOffset <= 1; yOffset++ {
				if xAdjusted, yAdjusted := x+xOffset, y+yOffset; isValidTile(board, xAdjusted, yAdjusted) {
					board.Field[xAdjusted][yAdjusted] += 1
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
	return nil
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
			if board.Field[x][y] == -1 {
				return -1
				// Returns -1 for mines to ensure correct behavior.
			}
			var minesFound = 0
			for xOffset := -1; xOffset <= 1; xOffset++ {
				for yOffset := -1; yOffset <= 1; yOffset++ {
					if xAdjusted, yAdjusted := x+xOffset, y+yOffset; isInRange(board, xAdjusted, yAdjusted) {
						if board.Field[xAdjusted][yAdjusted] == -1 {
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
		// Returns true if and only if all hints have the correct number of mines adjacent.
	}

	return mineCount() && hintVeracity()
}
