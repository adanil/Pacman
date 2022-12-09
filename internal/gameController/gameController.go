package gameController

import (
	"math"
	"pacman/internal/level"
	"pacman/internal/player"
)

type GameController struct {
	gameLevel *level.Level
	pacman    *player.Player
	TileSize  int
}

func InitGameController(gLevel *level.Level, pacman *player.Player, tileSize int) GameController {
	return GameController{gameLevel: gLevel, pacman: pacman, TileSize: tileSize}
}

func (c *GameController) UpdatePacman(direction int) {
	oldX := c.pacman.X
	oldY := c.pacman.Y
	oldRotation := c.pacman.Rotation
	c.pacman.ChangeDirection(direction)
	c.pacman.Move(direction)

	if c.CheckCollision(c.pacman.X, c.pacman.Y) {
		c.pacman.SetCoords(oldX, oldY)
		c.pacman.Rotation = oldRotation
		c.pacman.Move(oldRotation)

		if c.CheckCollision(c.pacman.X, c.pacman.Y) {
			c.pacman.SetCoords(oldX, oldY)
			c.pacman.Rotation = oldRotation
		}
	}
}

func (c *GameController) CheckCollision(newX, newY int) bool {
	xTileUp := int(math.Ceil(float64(c.pacman.X) / float64(c.TileSize)))
	yTileUp := int(math.Ceil(float64(c.pacman.Y) / float64(c.TileSize)))

	xTileDown := int(math.Floor(float64(c.pacman.X) / float64(c.TileSize)))
	yTileDown := int(math.Floor(float64(c.pacman.Y) / float64(c.TileSize)))

	return c.gameLevel.LevelTiles[xTileUp][yTileUp] == level.Wall || c.gameLevel.LevelTiles[xTileDown][yTileDown] == level.Wall || c.gameLevel.LevelTiles[xTileUp][yTileDown] == level.Wall || c.gameLevel.LevelTiles[xTileDown][yTileUp] == level.Wall
}
