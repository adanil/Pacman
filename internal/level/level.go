package level

import (
	"fmt"
	"math"
	"os"
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
	CreateLevel(width, height, tileSize int) Level
}

type Generator struct {
	Creator Creator
}

func (g *Generator) CreateLevel(width, height, tileSize int) Level {
	return g.Creator.CreateLevel(width, height, tileSize)
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

func (l *Level) CreateFood() int {
	foodCount := 0
	for x := 0; x < l.Width; x++ {
		for y := 0; y < l.Height; y++ {
			if l.LevelTiles[x][y] == Free {
				l.LevelTiles[x][y] = Food
				foodCount++
			}
		}
	}
	return foodCount - 1 //TODO CHECK IT
}

func (l *Level) UpdateAll() bool {
	l.UpdatePacman(&l.Player)
	for _, enemy := range l.Enemies {
		if l.UpdateEnemy(enemy) == false {
			return false
		}
	}
	return true
}
func (l *Level) UpdatePacman(player *entities.Pacman) {
	oldX, oldY := player.GetCoords()
	rotation := player.GetDirection()
	player.Move(rotation, l.Width*l.TileSize, l.Height*l.TileSize)

	//If a wall is encountered the coordinates do not change
	if l.CheckWallCollision(player.GetCoords()) {
		player.SetCoords(oldX, oldY)
		player.SetStopped(true)
		return
	} else {
		player.SetStopped(false)
	}

	xTileOld := (oldX + l.TileSize/2) / l.TileSize
	yTileOld := (oldY + l.TileSize/2) / l.TileSize
	l.LevelTiles[xTileOld%l.Width][yTileOld%l.Height] = Free

	newX, newY := player.GetCoords()
	xTileNew := (newX + l.TileSize/2) / l.TileSize
	yTileNew := (newY + l.TileSize/2) / l.TileSize

	if l.LevelTiles[xTileNew%l.Width][yTileNew%l.Height] == Food {
		player.Score++
	}
	l.LevelTiles[xTileNew%l.Width][yTileNew%l.Height] = Player

}

func (l *Level) UpdateEnemy(enemy entities.Playable) bool {
	oldX, oldY := enemy.GetCoords()
	rotation := enemy.GetDirection()
	enemy.Move(rotation, l.Width*l.TileSize, l.Height*l.TileSize)

	//If a wall is encountered the coordinates do not change
	if l.CheckWallCollision(enemy.GetCoords()) {
		enemy.SetCoords(oldX, oldY)
		enemy.SetStopped(true)
	} else {
		enemy.SetStopped(false)
	}
	if l.CheckHit(enemy.GetCoords()) {
		l.Player.Health--
		return false
	}
	return true
}

func (l *Level) CheckWallCollision(x, y int) bool {
	xTileUp := int(math.Ceil(float64(x) / float64(l.TileSize)))
	yTileUp := int(math.Ceil(float64(y) / float64(l.TileSize)))

	xTileDown := int(math.Floor(float64(x) / float64(l.TileSize)))
	yTileDown := int(math.Floor(float64(y) / float64(l.TileSize)))

	return l.LevelTiles[xTileUp%l.Width][yTileUp%l.Height] == Wall || l.LevelTiles[xTileDown%l.Width][yTileDown%l.Height] == Wall ||
		l.LevelTiles[xTileUp%l.Width][yTileDown%l.Height] == Wall || l.LevelTiles[xTileDown%l.Width][yTileUp%l.Height] == Wall
}

func (l *Level) CheckHit(x, y int) bool {
	pacmanX, pacmanY := l.Player.GetCoords()
	pacmanCenterX := (pacmanX + l.TileSize) / 2
	pacmanCenterY := (pacmanY + l.TileSize) / 2
	enemyCenterX := (x + l.TileSize) / 2
	enemyCenterY := (y + l.TileSize) / 2

	return math.Hypot(float64(enemyCenterX-pacmanCenterX), float64(enemyCenterY-pacmanCenterY)) < float64(l.TileSize)/3.0
}

func (l *Level) GameOver() {
	fmt.Println("Game over. Score: ", l.Player.Score)
	os.Exit(0)
}

func (l *Level) ReleaseDecorators() {

}
