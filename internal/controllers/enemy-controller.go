package controllers

import (
	"math/rand"
	"pacman/internal/command"
	"pacman/internal/entities"
	"pacman/internal/level"
)

type RandomEnemyController struct {
	level *level.Level
}

func NewRandomEnemyController(level_ *level.Level) RandomEnemyController {
	return RandomEnemyController{level: level_}
}

func (e *RandomEnemyController) GetCommands() []command.Command {
	var commands []command.Command
	for _, enemy := range e.level.Enemies() {
		if c := e.GetCommand(enemy); c != nil {
			commands = append(commands, c)
		}
	}
	return commands
}

func (e *RandomEnemyController) GetCommand(enemy *entities.Enemy) command.Command {
	x, y := enemy.GetCoords()
	if x%e.level.TileSize() != 0 || y%e.level.TileSize() != 0 {
		return nil
	}
	const directionsNumber = 4
	if n := rand.Intn(directionsNumber + 1); n != 4 && (n != entities.OppositeDirection[enemy.GetDirection()] || enemy.GetStopped()) {
		cdCommand := command.NewChangeDirectionCommand(n, enemy, e.level)
		return &cdCommand
	}
	return nil
}
