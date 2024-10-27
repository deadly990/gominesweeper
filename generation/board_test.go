package generation

import "testing"

func TestValidation_Functional(test *testing.T) {
	validBoard := new(Board)
	validBoard.Mines = 10
	validBoard.Field = [][]int{
		{1, 2, -1, -1, -1, -1, 1},
		{1, -1, 4, 4, 4, 2, 1},
		{1, 1, 3, -1, 2, 0, 0},
		{1, 1, 2, -1, 2, 0, 0},
		{-1, 1, 1, 1, 1, 0, 0},
		{2, 2, 1, 0, 0, 1, 1},
		{1, -1, 1, 0, 0, 1, -1},
	}
	validBoard.seed = 70

	if !Validate(*validBoard) {
		test.Fatalf("Expected Generation#Validate to return True. Actual: False")
	}
}

func TestValidation_AllMines(test *testing.T) {
	board := new(Board)
	board.Mines = 4
	board.Field = [][]int{
		{-1, -1},
		{-1, -1},
	}

	if validation := Validate(*board); !validation {
		test.Fatalf("Expected Generation#Validate to return True. Actual: %v", validation)
	}
}

func TestGeneration_ErrorOverload(test *testing.T) {
	if _, err := NewBoard(5, 2, 2, 1); err == nil {
		test.Fatalf("Expected Generation#NewBoard to produce MineOverload error. Actual: %v", err)
	}
}
