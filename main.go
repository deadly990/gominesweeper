package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/deadly990/gominesweeper/view"
)

var mainPageTemplate = view.Generate()

func main() {
	addr := flag.String("addr", ":80", "http service address")
	flag.Parse()
	http.Handle("/test", http.HandlerFunc(rootHandler))
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func rootHandler(w http.ResponseWriter, req *http.Request) {

	squares := make([][]view.Square, 5)
	for i := range squares {
		squares[i] = make([]view.Square, 5)
	}

	squares[1][2] = 3
	mineView := view.MineView{Squares: squares}
	mainData := view.MainData{Mine: mineView}
	err := mainPageTemplate.ExecuteTemplate(w, "mainpage.html", mainData)
	if err != nil {
		log.Fatal("ExecuteTemplate:", err)
	}
}
