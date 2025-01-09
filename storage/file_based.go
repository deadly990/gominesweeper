package storage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/deadly990/gominesweeper/game"
	"github.com/deadly990/gominesweeper/generation"
)

var cwd, _ = filepath.Abs(".")
var PathCrumb = filepath.Join(cwd, "saves")

type Move struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameSave struct {
	Seed      int64  `json:"seed"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	MineCount int    `json:"mineCount"`
	Moves     []Move `json:"moves"`
}

func FromGame(game game.Game) *GameSave {
	seed := game.Board.Seed
	width, height := generation.BoardSize(game.Board)
	mineCount := game.Board.Mines
	savedMoves := []Move{}
	for _, coordinate := range game.Moves {
		translation := &Move{coordinate.X, coordinate.Y}
		savedMoves = append(savedMoves, *translation)
	}
	return &GameSave{seed, width, height, mineCount, savedMoves}
}

func (gameSave *GameSave) ToGame() *game.Game {
	board, err := generation.NewBoard(
		gameSave.MineCount,
		gameSave.Width,
		gameSave.Height,
		gameSave.Seed,
	)
	if err != nil {
		log.Fatalf("Encountered an error in converting GameSave to Game: %s", err)
	}
	return game.NewGame(*board)
}

func (gameSave *GameSave) Save(name string) error {
	_, err := os.Stat(PathCrumb)
	if errors.Is(err, fs.ErrNotExist) {
		os.Mkdir(PathCrumb, 0755)
	}
	path := filepath.Join(PathCrumb, fmt.Sprintf("%s-%d.sweeper", name, gameSave.Seed))
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return gameSave.Encode(file)
}

func (gameSave *GameSave) Encode(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(gameSave)
}

func Load(name string) *GameSave {
	path := filepath.Join(PathCrumb, name+".sweeper")
	buffer, readErr := os.ReadFile(path)
	if readErr != nil {
		panic(readErr)
	}
	decoder := json.NewDecoder(bytes.NewReader(buffer))
	gameSave := GameSave{}
	decodeErr := decoder.Decode(&gameSave)
	if decodeErr != nil {
		panic(decodeErr)
	}
	return &gameSave
}
func (gameSave *GameSave) Decode(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(&gameSave)
}

func (receiver *Move) EquivalentTo(other Move) bool {
	return receiver.X == other.X && receiver.Y == other.Y
}

func (receiver *GameSave) EquivalentTo(other GameSave) bool {
	if receiver.Seed != other.Seed {
		return false
	}
	if receiver.Width != other.Width {
		return false
	}
	if receiver.Height != other.Height {
		return false
	}
	if receiver.MineCount != other.MineCount {
		return false
	}
	if len(receiver.Moves) != len(other.Moves) {
		return false
	}

	for index, move := range receiver.Moves {
		if !move.EquivalentTo(other.Moves[index]) {
			return false
		}
	}
	return true
}
