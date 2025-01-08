package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/deadly990/gominesweeper/game"
	"github.com/deadly990/gominesweeper/generation"
	"github.com/deadly990/gominesweeper/view"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var mainPageTemplate = view.Generate()

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/game", func(r chi.Router) {
		r.Get("/", rootHandler)
		r.Route("/{gameId}/click/{clickLocation}", func(r chi.Router) {
			r.Use(GameCtx)
			r.Use(ClickCtx)
		})
	})
	addr := flag.String("addr", ":80", "http service address")
	flag.Parse()
	http.Handle("/test", http.HandlerFunc(rootHandler))
	http.Handle("/click/:game/:name", http.HandlerFunc(clickHandler))
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	newBoard, boardErr := generation.NewBoard(25, 10, 10, time.Now().UnixNano())
	if boardErr != nil {
		log.Println("NewBoard:", boardErr)
		http.Error(w, boardErr.Error(), 500)
		return
	}
	game := game.NewGame(*newBoard)
	mineView := view.FromGame(*game)
	mainData := view.MainData{Mine: mineView}
	err := mainPageTemplate.ExecuteTemplate(w, "mainpage.html", mainData)
	if err != nil {
		log.Fatal("ExecuteTemplate:", err)
	}

}
func GameCtx(next http.Handler) http.Handler {
	return next
}
func ClickCtx(next http.Handler) http.Handler {
	return next
}
func clickHandler(w http.ResponseWriter, req *http.Request) {

}
