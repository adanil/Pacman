package controllers

import (
	"math/rand"
	"pacman/internal/command"
	"pacman/internal/level"
)

type RandomEnemyController struct {
	level      *level.Level
	updateFreq int
	//TODO add entity which controller controls
}

func NewRandomEnemyController(level_ *level.Level) RandomEnemyController {
	return RandomEnemyController{level: level_}
}

func (e *RandomEnemyController) GetCommands() []command.Command {
	var commands []command.Command
	for _, enemy := range e.level.Enemies {
		x, y := enemy.GetCoords()
		if x%e.level.TileSize != 0 || y%e.level.TileSize != 0 {
			continue
		}
		if n := rand.Intn(5); (n != 4 && e.updateFreq == 0) || (n != 4 && enemy.GetStopped()) {
			cdCommand := command.NewChangeDirectionCommand(n, enemy, e.level)
			commands = append(commands, &cdCommand)
		}
	}
	e.updateFreq++
	e.updateFreq %= 60
	return commands
}
