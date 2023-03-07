package controllers

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pacman/internal/command"
	"pacman/internal/entities"
	"pacman/internal/level"
)

type KeyboardHandler struct {
	pressedRotationButton int
	level                 *level.Level
}

func NewKeyboardHandler(level_ *level.Level) KeyboardHandler {
	return KeyboardHandler{pressedRotationButton: -1, level: level_}
}

func (k *KeyboardHandler) GetKeyboardCommands() []command.Command {
	var commands []command.Command

	if k.pressedRotationButton != -1 {
		cdCommand := command.NewChangeDirectionCommand(entities.Direction(k.pressedRotationButton), k.level.Player(), k.level)
		commands = append(commands, &cdCommand)
	}

	return commands
}

func (k *KeyboardHandler) HandlePressedButtons(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			switch {
			case inpututil.KeyPressDuration(ebiten.KeyArrowDown) > 0:
				k.pressedRotationButton = int(entities.DOWN)
			case inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0:
				k.pressedRotationButton = int(entities.UP)
			case inpututil.KeyPressDuration(ebiten.KeyArrowRight) > 0:
				k.pressedRotationButton = int(entities.RIGHT)
			case inpututil.KeyPressDuration(ebiten.KeyArrowLeft) > 0:
				k.pressedRotationButton = int(entities.LEFT)
			case inpututil.KeyPressDuration(ebiten.KeyS) > 0:
				k.pressedRotationButton = int(entities.DOWN)
			case inpututil.KeyPressDuration(ebiten.KeyW) > 0:
				k.pressedRotationButton = int(entities.UP)
			case inpututil.KeyPressDuration(ebiten.KeyD) > 0:
				k.pressedRotationButton = int(entities.RIGHT)
			case inpututil.KeyPressDuration(ebiten.KeyA) > 0:
				k.pressedRotationButton = int(entities.LEFT)
			default:
				k.pressedRotationButton = -1
			}
		}
	}
}
