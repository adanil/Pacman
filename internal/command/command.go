package command

import (
	"pacman/internal/entities"
	"pacman/internal/level"
)

type Command interface {
	Execute()
}

type ChangeDirectionCommand struct {
	direction int
	target    entities.Movable
	level     *level.Level
}

func NewChangeDirectionCommand(direction_ int, target_ entities.Movable, level_ *level.Level) ChangeDirectionCommand {
	return ChangeDirectionCommand{direction: direction_, target: target_, level: level_}
}

func (c *ChangeDirectionCommand) Execute() {
	oldX, oldY := c.target.GetCoords()
	oldRotation := c.target.GetDirection()
	c.target.ChangeDirection(c.direction)
	c.target.Move(c.direction, c.level.Width*c.level.TileSize, c.level.Height*c.level.TileSize) //TODO try to check possibility of change direction without using move method

	if c.level.CheckWallCollision(c.target.GetCoords()) {
		c.target.ChangeDirection(oldRotation)
	}

	c.target.SetCoords(oldX, oldY)

}
