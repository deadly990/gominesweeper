package view

import (
	"html/template"
)

type Square int

func visible(square Square) bool {
	return square > 0
}

type MineView struct {
	Time      int
	Remaining int
	Squares   [][]Square
}
type MainData struct {
	Mine MineView
}

func Generate() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"IsVisible": visible,
	}).ParseGlob("./templates/*"))
}
