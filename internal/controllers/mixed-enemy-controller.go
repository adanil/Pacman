package controllers

import (
	"math"
	"math/rand"
	"pacman/internal/command"
	"pacman/internal/level"
)

const maxThreshold = 100

type MixedEnemyController struct {
	level            *level.Level
	randomController RandomEnemyController
	spfController    SPFEnemyController
	randomThreshold  int
}

func NewMixedEnemyController(level_ *level.Level, randomThreshold int) MixedEnemyController {
	return MixedEnemyController{
		level:            level_,
		randomController: NewRandomEnemyController(level_),
		spfController:    NewSPFEnemyController(level_),
		randomThreshold:  int(math.Min(float64(randomThreshold), maxThreshold)),
	}
}

func (m *MixedEnemyController) GetCommands() []command.Command {
	var commands []command.Command
	for _, enemy := range m.level.Enemies() {
		n := rand.Intn(maxThreshold)
		var c command.Command
		if n > m.randomThreshold {
			c = m.spfController.GetCommand(enemy)
		} else {
			c = m.randomController.GetCommand(enemy)
		}
		if c != nil {
			commands = append(commands, c)
		}
	}
	return commands
}
