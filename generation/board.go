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

func getSeed(board Board) int64 {
	return board.seed
}
func BoardSize(board Board) (int, int) {
	return len(board.Field[0]), len(board.Field)
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
	valid, err := Validate(board)
	if !valid {
		return &board, fmt.Errorf("Board generation failed, board is invalid: %s", err)
	}
	return &board, genErr
}

func blankField(width int, height int) [][]int {
	var arr = make([][]int, height)
	for i := 0; i < width; i++ {
		arr[i] = make([]int, width)
	}
	return arr
}

func isInRange(board Board, y int, x int) bool {
	width, height := BoardSize(board)
	return (x >= 0 && x < width) && (y >= 0 && y < height)
}

// Returns false if the tile is out of bounds or is a mine.
func isValidTile(board Board, y int, x int) bool {
	return isInRange(board, y, x) && board.Field[y][x] != -9
}

func generateMines(board Board) error {
	width, height := BoardSize(board)

	var pregenTests = func() error {
		if board.Mines > width*height {
			return fmt.Errorf("board size specified cannot hold the number mines provided")
		}
		return nil // No error detected
	}
	if testResult := pregenTests(); testResult != nil {
		return testResult
	}

	// Traverses all adjacent tiles in a 1 tile radius
	var populateHints = func(y int, x int) {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			for xOffset := -1; xOffset <= 1; xOffset++ {
				if yAdjusted, xAdjusted := y+yOffset, x+xOffset; isValidTile(board, yAdjusted, xAdjusted) {
					board.Field[yAdjusted][xAdjusted] += 1
				}
			}
		}
	}

	var random = rand.New(rand.NewSource((board.seed)))
	// Iterates until n mines have been successfully placed.
	for count := 0; count < board.Mines; {
		var x = random.Intn(width)
		var y = random.Intn(height)
		if board.Field[y][x] == -9 {
			continue
			// Does not count to the progress of mines on the occasion that a mine already exists in a location.
		}
		count++
		board.Field[y][x] = -9
		populateHints(y, x)
	}
	return nil
}

func Validate(board Board) (bool, error) {
	width, height := BoardSize(board)

	var mineCount = func() int {
		var actual = 0
		for y := 0; y < width; y++ {
			for x := 0; x < height; x++ {
				if board.Field[y][x] == -9 {
					actual++
				}
			}
		}
		return actual
	}

	var hintVeracity = func() (bool, error) {
		var countSurroundings = func(y int, x int) int {
			if board.Field[y][x] == -9 {
				return -9
				// Returns -9 for mines to ensure correct behavior.
			}
			var minesFound = 0
			for yOffset := -1; yOffset <= 1; yOffset++ {
				for xOffset := -1; xOffset <= 1; xOffset++ {
					if yAdjusted, xAdjusted := y+yOffset, x+xOffset; isInRange(board, yAdjusted, xAdjusted) {
						if board.Field[yAdjusted][xAdjusted] == -9 {
							minesFound++
						}
					}
				}
			}
			return minesFound
		}

		for y := 0; y < width; y++ {
			for x := 0; x < height; x++ {
				if count := countSurroundings(y, x); board.Field[y][x] != count {
					return false, fmt.Errorf("hint differed from surrounding count hint %d, surroundingCount %d", board.Field[y][x], count)
				}
			}
		}
		return true, nil
		// Returns true if and only if all hints have the correct number of mines adjacent.
	}
	actual := mineCount()
	if actual > board.Mines {
		return false, fmt.Errorf("too many mines placed Actual %d, Expected %d", actual, board.Mines)
	} else if actual < board.Mines {
		return false, fmt.Errorf("too few mines placed Actual %d, Expected %d", actual, board.Mines)
	}
	return hintVeracity()
}
