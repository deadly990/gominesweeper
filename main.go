package main

import (
	"fmt"

	"github.com/deadly990/gominesweeper/generation"
)

func main() {
	var board = generation.NewBoard(10, 7, 7, 70)
	fmt.Println(board.Field)
	fmt.Print("Board validation: ")
	fmt.Println(generation.Validate(*board))
}
