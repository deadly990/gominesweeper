package controllers

import (
	"github.com/deadly990/gominesweeper/game"
)

type ClickType string

const (
	leftClick   ClickType = "left"
	rightClick  ClickType = "right"
	middleClick ClickType = "middle"
)

type ClickCommand struct {
	Type        ClickType
	YCoordinate int
	XCoordinate int
}

func RunClickCommand(gameInstance game.Game, command ClickCommand) game.Game {
	// switch command.Type {
	// case leftClick:
	// case rightClick:
	// case middleClick:
	// }
	gameInstance.Move(game.Coordinate{X: command.XCoordinate, Y: command.YCoordinate}, gameInstance.Clear)
	return gameInstance
}
