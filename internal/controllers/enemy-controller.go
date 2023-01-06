package controllers

import (
	"math/rand"
	"pacman/internal/command"
	"pacman/internal/level"
)

type EnemyController struct {
	level *level.Level
	//TODO add entity which controller controls
}

func NewEnemyController(level_ *level.Level) EnemyController {
	return EnemyController{level: level_}
}

func (e *EnemyController) GetCommands() []command.Command {
	var commands []command.Command
	for _, enemy := range e.level.Enemies {
		if n := rand.Intn(5); n != 4 {
			cdCommand := command.NewChangeDirectionCommand(n, enemy, e.level)
			commands = append(commands, &cdCommand)
		}
	}
	return commands
}
