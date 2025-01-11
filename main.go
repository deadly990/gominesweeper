package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/deadly990/gominesweeper/game"
	"github.com/deadly990/gominesweeper/generation"
	"github.com/deadly990/gominesweeper/storage"
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
	r.Get("/test", http.HandlerFunc(rootHandler))
	addr := flag.String("addr", ":80", "http service address")
	flag.Parse()
	// http.Handle("/test", http.HandlerFunc(rootHandler))
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	newBoard, boardErr := generation.NewBoard(25, 10, 10, int64(random.Int63()))
	if boardErr != nil {
		log.Println("NewBoard:", boardErr)
		http.Error(w, boardErr.Error(), 500)
		return
	}
	game := game.NewGame(*newBoard)

	gameName := generateName(rand.Int63())
	mineView := view.FromGame(*game, gameName)
	mainData := view.MainData{Mine: mineView}
	err := mainPageTemplate.ExecuteTemplate(w, "mainpage.html", mainData)
	if err != nil {
		log.Fatal("ExecuteTemplate:", err)
	}
	storage.FromGame(*game).Save(gameName)
}

func generateName(seed int64) string {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(seed))
	hash := sha256.Sum256(buf)
	return hex.EncodeToString(hash[:])
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
	// Handle click
	gameCtx := req.Context().Value(GameIDString).(string)
	clickCtx := req.Context().Value(ClickLocationString).(string)

	coord, err := parseClickLocation(clickCtx)
	if err != nil {
		log.Fatal("parseClickLocation:", err)
	}
	game := storage.Load(gameCtx).ToGame()
	game.Move(coord, game.Clear)

	storage.FromGame(*game).Save(gameCtx)

	// Display updated board
	mineView := view.FromGame(*game, gameCtx)
	mainData := view.MainData{Mine: mineView}
	err = mainPageTemplate.ExecuteTemplate(w, "mainpage.html", mainData)
	if err != nil {
		log.Fatal("ExecuteTemplate:", err)
	}

	log.Printf("Game: %s Click: %s", gameCtx, clickCtx)
}

func parseClickLocation(clickCtx string) (game.Coordinate, error) {
	clickArr := strings.Split(clickCtx, "_")
	if len(clickArr) != 2 {
		return game.Coordinate{}, fmt.Errorf("click location parsing encountered an error. Array length was unexpected. Actual: %+v", clickArr)
	}
	intArr := []int{}
	for _, val := range clickArr {
		converted, err := strconv.Atoi(val)
		if err != nil {
			return game.Coordinate{}, fmt.Errorf("click location parsing encountered an error. Error occurred whilst parsing string to integer. Actual: %s", val)
		}
		intArr = append(intArr, converted)
	}
	return game.Coordinate{X: intArr[1], Y: intArr[0]}, nil
}
