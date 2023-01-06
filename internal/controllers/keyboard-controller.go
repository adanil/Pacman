package controllers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pacman/internal/command"
	"pacman/internal/entities"
	"pacman/internal/level"
)

type KeyboardHandler struct {
	pressedRotationButton int
	level                 *level.Level
	//TODO add entity which controller controls
}

func NewKeyboardHandler(level_ *level.Level) KeyboardHandler {
	return KeyboardHandler{pressedRotationButton: -1, level: level_}
}

func (k *KeyboardHandler) GetKeyboardCommands() []command.Command {
	var commands []command.Command

	if k.pressedRotationButton != -1 {
		cdCommand := command.NewChangeDirectionCommand(k.pressedRotationButton, &k.level.Player, k.level)
		commands = append(commands, &cdCommand)
	}

	return commands
}

func (k *KeyboardHandler) HandlePressedButtons() {
	for {
		if inpututil.KeyPressDuration(ebiten.KeyArrowDown) > 0 {
			k.pressedRotationButton = entities.DOWN
		} else if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
			k.pressedRotationButton = entities.UP
		} else if inpututil.KeyPressDuration(ebiten.KeyArrowRight) > 0 {
			k.pressedRotationButton = entities.RIGHT
		} else if inpututil.KeyPressDuration(ebiten.KeyArrowLeft) > 0 {
			k.pressedRotationButton = entities.LEFT
		} else {
			k.pressedRotationButton = -1
		}
	}

}
