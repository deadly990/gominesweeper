package generation

import (
	"testing"
)

func TestValidation(test *testing.T) {
	var validBoard Board
	validBoard.Mines = 10
	validBoard.Field = [][]int{
		{1, 2, -1, -1, -1, -1, 1},
		{1, -1, 4, 4, 4, 2, 1},
		{1, 1, 3, -1, 2, 0, 0},
		{1, 1, 2, -1, 2, 0, 0},
		{-1, 1, 1, 1, 1, 0, 0},
		{2, 2, 1, 0, 0, 1, 1},
		{1, -1, 1, 0, 0, 1, -1}}
	validBoard.seed = 70

	if Validate(validBoard) == false {
		test.Fatalf("Expected #Validate to return True. Actual return: False")
	}

}
