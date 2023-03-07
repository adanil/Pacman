package command

import (
	"pacman/internal/base"
	"pacman/internal/entities"
	"pacman/internal/level"
)

type Command interface {
	Execute()
}

type ChangeDirectionCommand struct {
	direction entities.Direction
	target    entities.Movable
	level     *level.Level
}

func NewChangeDirectionCommand(direction_ entities.Direction, target_ entities.Movable, level_ *level.Level) ChangeDirectionCommand {
	return ChangeDirectionCommand{direction: direction_, target: target_, level: level_}
}

func (c *ChangeDirectionCommand) Execute() {
	nextX, nextY := c.target.CalculateNextPosition(c.direction, base.GameScreenWidth, base.GameScreenHeight)
	if !c.level.CheckWallCollision(nextX, nextY) {
		c.target.ChangeDirection(c.direction)
	}
}
