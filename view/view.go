package view

import (
	"fmt"
	"html/template"

	"github.com/deadly990/gominesweeper/game"
	"github.com/deadly990/gominesweeper/generation"
)

type Square struct {
	Value int
	Name  string
}

func visible(square Square) bool {
	return square.Value > 0
}

type MineView struct {
	Remaining int
	Squares   [][]Square
}
type MainData struct {
	Mine MineView
}

func convert(field [][]int) [][]Square {
	squares := make([][]Square, len(field))
	for i := range squares {
		squares[i] = make([]Square, len(field[i]))
	}
	for i := range field {
		for j := range field[i] {
			squares[i][j] = Square{
				Value: field[i][j],
				Name:  fmt.Sprintf("%d_%d", i, j),
			}
		}
	}
	return squares
}
func FromBoard(board generation.Board) MineView {
	return MineView{
		Remaining: board.Mines,
		Squares:   convert(board.Field),
	}
}
func FromGame(game game.Game) MineView {
	return MineView{
		Remaining: game.Board.Mines,
		Squares:   convert(game.Revealed),
	}
}
func Generate() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"IsVisible": visible,
	}).ParseGlob("./templates/*"))
}
