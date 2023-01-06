package level

import (
	"math"
	"pacman/internal/entities"
)

const (
	Free = iota
	Wall
	Player
	Bot
	Food
)

type Creator interface {
	CreateLevel(width, height, tileSize, enemyCount int) Level
}

type Generator struct {
	Creator Creator
}

func (g *Generator) CreateLevel(width, height, tileSize, enemyCount int) Level {
	return g.Creator.CreateLevel(width, height, tileSize, enemyCount)
}

type Level struct {
	LevelTiles     [][]int
	TileSize       int
	Width          int
	Height         int
	Player         entities.Pacman
	Enemies        []entities.Playable
	DecoratorTimer map[int]entities.Playable
}

func (l *Level) CreateFood() {
	for x := 0; x < l.Width; x++ {
		for y := 0; y < l.Height; y++ {
			if l.LevelTiles[x][y] == Free {
				l.LevelTiles[x][y] = Food
			}
		}
	}
}

func (l *Level) UpdateAll() {
	l.UpdatePacman(&l.Player)
	for _, enemy := range l.Enemies {
		l.UpdateEnemy(enemy)
	}
}
func (l *Level) UpdatePacman(player *entities.Pacman) {
	oldX, oldY := player.GetCoords()
	rotation := player.GetDirection()
	player.Move(rotation)

	//If a wall is encountered the coordinates do not change
	if l.CheckWallCollision(player.GetCoords()) {
		player.SetCoords(oldX, oldY)
		return
	}

	xTileOld := (oldX + l.TileSize/2) / l.TileSize
	yTileOld := (oldY + l.TileSize/2) / l.TileSize
	l.LevelTiles[xTileOld][yTileOld] = Free

	newX, newY := player.GetCoords()
	xTileNew := (newX + l.TileSize/2) / l.TileSize
	yTileNew := (newY + l.TileSize/2) / l.TileSize

	if l.LevelTiles[xTileNew][yTileNew] == Food {
		player.Score++
	}
	l.LevelTiles[xTileNew][yTileNew] = Player

}

func (l *Level) UpdateEnemy(enemy entities.Playable) {
	oldX, oldY := enemy.GetCoords()
	rotation := enemy.GetDirection()
	enemy.Move(rotation)

	//If a wall is encountered the coordinates do not change
	if l.CheckWallCollision(enemy.GetCoords()) {
		enemy.SetCoords(oldX, oldY)
		return
	}

}

func (l *Level) CheckWallCollision(x, y int) bool {
	xTileUp := int(math.Ceil(float64(x) / float64(l.TileSize)))
	yTileUp := int(math.Ceil(float64(y) / float64(l.TileSize)))

	xTileDown := int(math.Floor(float64(x) / float64(l.TileSize)))
	yTileDown := int(math.Floor(float64(y) / float64(l.TileSize)))

	return l.LevelTiles[xTileUp][yTileUp] == Wall || l.LevelTiles[xTileDown][yTileDown] == Wall ||
		l.LevelTiles[xTileUp][yTileDown] == Wall || l.LevelTiles[xTileDown][yTileUp] == Wall
}

func (l *Level) ReleaseDecorators() {

}
