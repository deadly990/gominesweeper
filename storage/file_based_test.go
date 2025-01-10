package storage

import (
	"bytes"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/deadly990/gominesweeper/game"
	"github.com/deadly990/gominesweeper/generation"
)

func TestEncoding(test *testing.T) {
	var board generation.Board
	board.Field = [][]int{
		{2, 2, 1, 0},   // [ 2,  2,  1, 0]
		{-9, -9, 1, 0}, // [-9, -9,  1, 0]
		{-9, 4, 2, 1},  // [-9,  4,  2, 1]
		{1, 2, -9, 1},  // [ 1,  2, -9, 1]
		{0, 1, 1, 1},   // [ 0,  1,  1, 1]
	}
	game := game.NewGame(board)
	game.Clear(0, 3)
	gameSave := FromGame(*game)
	buf := new(bytes.Buffer)
	err := gameSave.Encode(buf)
	if err != nil {
		log.Printf("Error in GameSave encoding: %s", err)
		test.FailNow()
	}

	result := buf.String()
	if result != `{"seed":0,"width":4,"height":5,"mineCount":0,"moves":[{"x":3,"y":0}]}`+"\n" { // JSON Encoding adds a newline after encoding. Added \n to expect correct result.
		log.Printf("GameSave encoding did not produce expected result. Actual: %v", result)
		test.Fail()
	}
}

func TestDecoding(test *testing.T) {
	decoded := &GameSave{}
	reader := strings.NewReader(`{"seed":0,"width":4,"height":5,"mineCount":0,"moves":[{"x":3,"y":0}]}`)
	err := decoded.Decode(reader)
	if err != nil {
		log.Printf("Decoding from JSON String produced an error: %s", err)
		test.FailNow()
	}
	move := Move{3, 0}
	expected := &GameSave{0, 4, 5, 0, []Move{move}}
	if !expected.EquivalentTo(*decoded) {
		log.Printf("Decoded GameSave did not produce expected results. Actual: %+v", decoded)
		test.Fail()
	}
}

func TestCodecInverseEquivalence(test *testing.T) {
	for iteration := 0; iteration < 100000; iteration++ {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		seed := random.Int63()
		width := random.Intn(23) + 2
		height := random.Intn(23) + 2
		mineCount := random.Intn(width*height/2) + 1

		board, genErr := generation.NewBoard(mineCount, width, height, seed)
		if genErr != nil {
			log.Printf("Test failed due to board generation error: %s", genErr)
			log.Printf("Failing criteria: mineCount:%d width:%d height:%d seed:%d", mineCount, width, height, seed)
			test.FailNow()
		}

		newGame := game.NewGame(*board)
		gameSave := FromGame(*newGame)
		buf := new(bytes.Buffer)
		encodeErr := gameSave.Encode(buf)
		if encodeErr != nil {
			log.Printf("Error in GameSave encoding: %s", encodeErr)
			log.Printf("Failing criteria: mineCount:%d width:%d height:%d seed:%d", mineCount, width, height, seed)
			test.FailNow()
		}
		decodedSave := &GameSave{}
		decodedSave.Decode(buf)

		if !gameSave.EquivalentTo(*decodedSave) {
			log.Printf("Decoded GameSave was not equivalent to its encoded self.")
			log.Printf("Failing criteria: mineCount:%d width:%d height:%d seed:%d", mineCount, width, height, seed)
			test.FailNow()
		}
	}
	// Produce seed, width, height, and mine count of failed test.
}
