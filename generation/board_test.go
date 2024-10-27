package generation

import (
	"fmt"
	"log"
	"testing"
)

func TestValidation_Functional(test *testing.T) {
	validBoard := new(Board)
	validBoard.Mines = 10
	validBoard.Field = [][]int{
		{1, 2, -9, -9, -9, -9, 1},
		{1, -9, 4, 4, 4, 2, 1},
		{1, 1, 3, -9, 2, 0, 0},
		{1, 1, 2, -9, 2, 0, 0},
		{-9, 1, 1, 1, 1, 0, 0},
		{2, 2, 1, 0, 0, 1, 1},
		{1, -9, 1, 0, 0, 1, -9},
	}
	validBoard.seed = 70

	if validation, err := Validate(*validBoard); !validation {
		log.Printf("Expected Generation#Validate to return True. Actual: %v Error: %s\n", validation, err)
		test.Fail()
	}
}

func TestValidation_AllMines(test *testing.T) {
	board := new(Board)
	board.Mines = 4
	board.Field = [][]int{
		{-9, -9},
		{-9, -9},
	}

	if validation, err := Validate(*board); !validation {
		log.Printf("Expected Generation#Validate to return True. Actual: %v Error: %s\n", validation, err)
		test.Fail()
	}
}

func TestGeneration_ErrorOverload(test *testing.T) {
	if _, err := NewBoard(5, 2, 2, 1); err == nil {
		fmt.Printf("Expected Generation#NewBoard to produce error. Actual: %v\n", err)
		test.Fail()
	}
}
