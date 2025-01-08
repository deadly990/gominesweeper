package main

import (
	"context"
	"flag"
	"fmt"
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

type contextName string

const GameIDString contextName = "gameId"
const ClickLocationString contextName = "clickLocation"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/game", func(r chi.Router) {
		r.Get("/", rootHandler)
		r.Route(fmt.Sprintf("/{%s}/click/{%s}", GameIDString, ClickLocationString), func(r chi.Router) {
			r.Use(GameCtx)
			r.Use(ClickCtx)
			r.Get("/", clickHandler)
		})
	})
	addr := flag.String("addr", ":80", "http service address")
	flag.Parse()
	http.Handle("/test", http.HandlerFunc(rootHandler))
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(*addr, r)
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
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		game := chi.URLParam(req, string(GameIDString))
		ctx := context.WithValue(req.Context(), GameIDString, game)
		next.ServeHTTP(w, req.WithContext(ctx))
	})

}
func ClickCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		click := chi.URLParam(req, string(ClickLocationString))
		ctx := context.WithValue(req.Context(), ClickLocationString, click)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
func clickHandler(w http.ResponseWriter, req *http.Request) {
	gameCtx := req.Context().Value(GameIDString)
	clickCtx := req.Context().Value(ClickLocationString)

	log.Default().Printf("Game: %s Click: %s", gameCtx, clickCtx)
}
