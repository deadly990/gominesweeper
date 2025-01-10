package view

import (
	"fmt"
	"html/template"

	"github.com/deadly990/gominesweeper/game"
	"github.com/deadly990/gominesweeper/generation"
)

type Tile struct {
	Value    int
	Location string
	GameID   string
}

func visible(square Tile) bool {
	return square.Value > 0
}

type MineView struct {
	Remaining int
	Squares   [][]Tile
	Name      string
}
type MainData struct {
	Mine MineView
}

func convert(field [][]int, game string) [][]Tile {
	squares := make([][]Tile, len(field))
	for i := range squares {
		squares[i] = make([]Tile, len(field[i]))
	}
	for i := range field {
		for j := range field[i] {
			squares[i][j] = Tile{
				Value:    field[i][j],
				Location: fmt.Sprintf("%d_%d", i, j),
				GameID:   game,
			}
		}
	}
	return squares
}
func FromBoard(board generation.Board, name string) MineView {
	return MineView{
		Remaining: board.Mines,
		Squares:   convert(board.Field, name),
	}
}
func FromGame(game game.Game, name string) MineView {
	return MineView{
		Remaining: game.Board.Mines,
		Squares:   convert(game.Revealed, name),
	}
}
func Generate() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"IsVisible": visible,
	}).ParseGlob("./templates/*"))
}
