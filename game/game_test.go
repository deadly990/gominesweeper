package game

import (
	"log"
	"testing"

	"github.com/deadly990/gominesweeper/generation"
)

func areArraysEqual(arr1 [][]int, arr2 [][]int) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[0]); j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}
	return true
}

func TestClearProliferate(test *testing.T) {
	var board generation.Board

	board.Field = [][]int{
		{0, 1, 2, 4, -9, 3, -9, 1, 1, -9},   // [ 0,  1,  2,  4, -9,  3, -9,  1,  1, -9]
		{0, 1, -9, -9, -9, 4, 1, 1, 2, 2},   // [ 0,  1, -9, -9, -9,  4,  1,  1,  2,  2]
		{0, 1, 2, 4, -9, 3, 2, 1, 2, -9},    // [ 0,  1,  2,  4, -9,  3,  2,  1,  2, -9]
		{2, 2, 1, 1, 2, -9, 2, -9, 4, 3},    // [ 2,  2,  1,  1,  2, -9,  2, -9,  4,  3]
		{-9, -9, 2, 0, 2, 3, 4, 3, -9, -9},  // [-9, -9,  2, {0}, 2,  3,  4,  3, -9, -9]
		{3, -9, 2, 0, 1, -9, -9, 2, 2, 2},   // [ 3, -9,  2,  0,  1, -9, -9,  2,  2,  2]
		{2, 2, 2, 0, 1, 2, 2, 1, 1, 1},      // [ 2,  2,  2,  0,  1,  2,  2,  1,  1,  1]
		{1, -9, 1, 0, 0, 0, 0, 0, 2, -9},    // [ 1, -9,  1,  0,  0,  0,  0,  0,  2, -9]
		{3, 3, 3, 1, 2, 2, 2, 1, 2, -9},     // [ 3,  3,  3,  1,  2,  2,  2,  1,  2, -9]
		{-9, -9, 2, -9, 2, -9, -9, 1, 1, 1}, // [-9, -9,  2, -9,  2, -9, -9,  1,  1,  1]
	}

	expected := [][]int{
		{-10, -1, -2, -4, -9, -3, -9, -1, -1, -9}, // [-10, -1, -2, -4, -9, -3, -9, -1, -1, -9]
		{-10, -1, -9, -9, -9, -4, -1, -1, -2, -2}, // [-10, -1, -9, -9, -9, -4, -1, -1, -2, -2]
		{-10, -1, -2, -4, -9, -3, -2, -1, -2, -9}, // [-10, -1, -2, -4, -9, -3, -2, -1, -2, -9]
		{-2, -2, 1, 1, 2, -9, -2, -9, -4, -3},     // [ -2, -2,  1,  1,  2, -9, -2, -9, -4, -3]
		{-9, -9, 2, 0, 2, -3, -4, -3, -9, -9},     // [ -9, -9,  2,  0,  2, -3, -4, -3, -9, -9]
		{-3, -9, 2, 0, 1, -9, -9, -2, -2, -2},     // [ -3, -9,  2,  0,  1, -9, -9, -2, -2, -2]
		{-2, -2, 2, 0, 1, 2, 2, 1, 1, -1},         // [ -2, -2,  2,  0,  1,  2,  2,  1,  1, -1]
		{-1, -9, 1, 0, 0, 0, 0, 0, 2, -9},         // [ -1, -9,  1,  0,  0,  0,  0,  0,  2, -9]
		{-3, -3, 3, 1, 2, 2, 2, 1, 2, -9},         // [ -3, -3,  3,  1,  2,  2,  2,  1,  2, -9]
		{-9, -9, -2, -9, -2, -9, -9, -1, -1, -1},  // [ -9, -9, -2, -9, -2, -9, -9, -1, -1, -1]
	}

	game := *NewGame(board)
	Move(game, 4, 3)
	if !areArraysEqual(game.Revealed, expected) {
		log.Printf("Move did not clear all blank tiles in the move area and reveal adjacent hints")
		test.Fail()
	}
}

// Comment
func TestDiagonalCornerReveal(test *testing.T) {
	var board generation.Board
	board.Field = [][]int{
		{2, 2, 1, 0},   // [ 2,  2,  1, 0]
		{-9, -9, 1, 0}, // [-9, -9,  1, 0]
		{-9, 4, 2, 1},  // [-9,  4,  2, 1]
		{1, 2, -9, 1},  // [ 1,  2, -9, 1]
		{0, 1, 1, 1},   // [ 0,  1,  1, 1]
	}

	expected := [][]int{
		{-2, -2, -1, -10}, // [-2, -2, -1, -0]
		{-9, -9, -1, -10}, // [-9, -9, -1, -0]
		{-9, -4, -2, -1},  // [-9, -4, -2, -1]
		{1, 2, -9, -1},    // [ 1,  2, -9, -1]
		{0, 1, -1, -1},    // [{0}, 1, -1, -1]
	}

	game := *NewGame(board)
	Move(game, 4, 0)

	if !areArraysEqual(game.Revealed, expected) {
		log.Printf("Move did not go as expected.")
		test.Fail()
	}
}

func TestOrderOfOperations(test *testing.T) {
	board, _ := generation.NewBoard(50, 100, 100, 1)

	game1 := *NewGame(*board)

	game2 := *NewGame(*board)

	Move(game1, 0, 0)
	Move(game1, 1, 1)
	Move(game1, 50, 50)

	Move(game2, 50, 50)
	Move(game2, 1, 1)
	Move(game2, 0, 0)

	if areArraysEqual(game1.Revealed, game2.Revealed) {
		log.Printf("move order changes revealed")
		test.Fail()
	}
}
