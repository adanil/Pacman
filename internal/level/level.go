package level

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
	"pacman/internal/base"
	"pacman/internal/entities"
	"time"
)

const (
	Free = iota
	Wall
	Player
	Food
	Strawberry
	NightModeBooster
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

type ScreenText struct {
	X, Y        int
	Text        string
	ExpiredTime time.Time
}

type Level struct {
	LevelTiles     [][]int
	TileSize       int
	Width          int
	Height         int
	Player         entities.Pacman
	Enemies        []*entities.Enemy
	DecoratorTimer map[int]entities.Entity
	Score          int
	FoodEaten      int
	Texts          []ScreenText
	FoodCount      int
}

func (l *Level) CreateEntities() {
	l.Player = l.CreateRandomPlayer(base.Images["pacman"])
	blueEnemy := entities.CreateEnemy(12*l.TileSize, 9*l.TileSize, base.Images["blueEnemy"])
	pinkEnemy := entities.CreateEnemy(11*l.TileSize, 9*l.TileSize, base.Images["pinkEnemy"])
	redEnemy := entities.CreateEnemy(10*l.TileSize, 9*l.TileSize, base.Images["redEnemy"])
	yellowEnemy := entities.CreateEnemy(9*l.TileSize, 9*l.TileSize, base.Images["yellowEnemy"])

	l.Enemies = append(l.Enemies, &blueEnemy, &pinkEnemy, &redEnemy, &yellowEnemy)

	l.CreateNightModeBoosters()
	l.CreateStrawberry()
	l.CreateFood()
}

func (l *Level) CreateRandomPlayer(playerImage *ebiten.Image) entities.Pacman {
	for {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(l.Width)
		y := rand.Intn(l.Height)
		if l.LevelTiles[x][y] == Free && !isEnemySpawn(Coordinates{x, y}) {
			p := entities.CreatePacman(x*l.TileSize, y*l.TileSize, playerImage)
			return p
		}
	}
}

func (l *Level) CreateFood() {
	for x := 0; x < l.Width; x++ {
		for y := 0; y < l.Height; y++ {
			if l.LevelTiles[x][y] == Free && !isEnemySpawn(Coordinates{x, y}) {
				l.LevelTiles[x][y] = Food
				l.FoodCount++
			}
		}
	}
}

func (l *Level) CreateStrawberry() {
	rand.Seed(time.Now().UnixNano())
	countStrawberry := rand.Intn(3) + 2
	for countStrawberry > 0 {
		var x, y int
		for {
			x = rand.Intn(l.Width)
			y = rand.Intn(l.Height)
			if l.LevelTiles[x][y] == Free && !isEnemySpawn(Coordinates{x, y}) {
				break
			}
		}
		l.LevelTiles[x][y] = Strawberry
		countStrawberry--
	}
}
func (l *Level) CreateNightModeBoosters() {
	rand.Seed(time.Now().UnixNano())
	countBoosters := rand.Intn(8) + 2
	for countBoosters > 0 {
		var x, y int
		for {
			x = rand.Intn(l.Width)
			y = rand.Intn(l.Height)
			if l.LevelTiles[x][y] == Free && !isEnemySpawn(Coordinates{x, y}) {
				break
			}
		}
		l.LevelTiles[x][y] = NightModeBooster
		countBoosters--
	}
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
	rotation := player.GetDirection()
	nextX, nextY := player.CalculateNextPosition(rotation, l.Width*l.TileSize, l.Height*l.TileSize)

	//If a wall is encountered the coordinates do not change
	x, y := player.GetCoords()
	xTileTmp := x / l.TileSize
	yTileTmp := y / l.TileSize
	if (l.CheckWallCollision(nextX, nextY)) || (xTileTmp == 10 && yTileTmp == 7 && player.GetDirection() == entities.DOWN) || (xTileTmp == 11 && yTileTmp == 7 && player.GetDirection() == entities.DOWN) {
		return
	}
	player.Move(rotation, l.Width*l.TileSize, l.Height*l.TileSize)

	xTile := ((nextX + l.TileSize/2) / l.TileSize) % l.Width
	yTile := ((nextY + l.TileSize/2) / l.TileSize) % l.Height

	object := l.LevelTiles[xTile][yTile]
	switch object {
	case Food:
		l.Score++
		l.FoodEaten++
	case Strawberry:
		l.Score += 200
		l.Texts = append(l.Texts, ScreenText{
			X:           xTile,
			Y:           yTile,
			Text:        "+200",
			ExpiredTime: time.Now().Add(3 * time.Second),
		})
	case NightModeBooster:
		for _, enemy := range l.Enemies {
			enemy.SetNightMode(true)
		}
	}

	l.LevelTiles[xTile][yTile] = Free
}

func (l *Level) UpdateEnemy(enemy *entities.Enemy) bool {
	if enemy.NightMode() && enemy.NightModeExpiredTime().Before(time.Now()) {
		enemy.SetNightMode(false)
	}

	rotation := enemy.GetDirection()
	nextX, nextY := enemy.CalculateNextPosition(rotation, l.Width*l.TileSize, l.Height*l.TileSize)

	//If a wall is encountered the coordinates do not change
	if l.CheckWallCollision(nextX, nextY) {
		enemy.SetStopped(true)
	} else {
		enemy.Move(rotation, l.Width*l.TileSize, l.Height*l.TileSize)
		enemy.SetStopped(false)
	}

	if l.CheckHit(enemy.GetCoords()) {
		if !enemy.NightMode() {
			l.Player.DecreaseHealth()
			return false
		}
		l.enemyKilled(enemy)
	}

	return true
}

func (l *Level) enemyKilled(enemy *entities.Enemy) {
	x, y := enemy.GetCoords()
	xTile := x / l.TileSize
	yTile := y / l.TileSize
	l.Score += 400
	l.Texts = append(l.Texts, ScreenText{
		X:           xTile,
		Y:           yTile,
		Text:        "+400",
		ExpiredTime: time.Now().Add(3 * time.Second),
	})
	//Respawn enemy
	enemy.SetCoords(enemy.GetStartCoords())
	enemy.SetNightMode(false)
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

func (l *Level) IsAllFoodEaten() bool {
	return l.FoodCount == l.FoodEaten
}