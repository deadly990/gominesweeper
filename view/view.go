package view

import (
	"html/template"

	"github.com/deadly990/gominesweeper/generation"
)

type Square int

func visible(square Square) bool {
	return square > 0
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
			squares[i][j] = Square(field[i][j])
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
func Generate() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"IsVisible": visible,
	}).ParseGlob("./templates/*"))
}
