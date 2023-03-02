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

const textDuration = 3
const strawberryMax = 5
const strawberryMin = 2

type Level struct {
	levelTiles [][]int
	tileSize   int
	width      int
	height     int
	player     entities.Pacman
	enemies    []*entities.Enemy
	score      int
	foodEaten  int
	texts      []ScreenText
	foodCount  int
}

func NewLevel(width, height, tileSize int) Level {
	return Level{levelTiles: make([][]int, width), width: width, height: height, tileSize: tileSize}
}

func (l *Level) CreateEntities() {
	l.player = l.createRandomPlayer(base.Images["pacman"])

	x, y := enemiesSpawnCoords[0].ToPixels()
	blueEnemy := entities.CreateEnemy(x, y, base.Images["blueEnemy"])

	x, y = enemiesSpawnCoords[1].ToPixels()
	pinkEnemy := entities.CreateEnemy(x, y, base.Images["pinkEnemy"])

	x, y = enemiesSpawnCoords[2].ToPixels()
	redEnemy := entities.CreateEnemy(x, y, base.Images["redEnemy"])

	x, y = enemiesSpawnCoords[3].ToPixels()
	yellowEnemy := entities.CreateEnemy(x, y, base.Images["yellowEnemy"])

	l.enemies = append(l.enemies, &blueEnemy, &pinkEnemy, &redEnemy, &yellowEnemy)

	l.createNightModeBoosters()
	l.createStrawberry()
	l.createFood()
}

func (l *Level) UpdateAll() bool {
	l.UpdatePacman(&l.player)
	for _, enemy := range l.enemies {
		if !l.UpdateEnemy(enemy) {
			return false
		}
	}
	return true
}

func (l *Level) UpdatePacman(player *entities.Pacman) {
	rotation := player.GetDirection()
	nextX, nextY := player.CalculateNextPosition(rotation, base.GameScreenWidth, base.GameScreenHeight)

	// If a wall is encountered the coordinates do not change
	x, y := player.GetCoords()
	xTileTmp := x / l.tileSize
	yTileTmp := y / l.tileSize
	if (l.CheckWallCollision(nextX, nextY)) || (xTileTmp == 10 && yTileTmp == 7 && player.GetDirection() == entities.DOWN) || (xTileTmp == 11 && yTileTmp == 7 && player.GetDirection() == entities.DOWN) {
		return
	}
	player.Move(rotation, base.GameScreenWidth, base.GameScreenWidth)

	xTile := ((nextX + l.tileSize/2) / l.tileSize) % l.width
	yTile := ((nextY + l.tileSize/2) / l.tileSize) % l.height

	object := l.levelTiles[xTile][yTile]
	switch object {
	case Food:
		l.score++
		l.foodEaten++
	case Strawberry:
		l.score += 200
		l.texts = append(l.texts, NewScreenText(xTile, yTile, "+200", time.Now().Add(textDuration*time.Second)))
	case NightModeBooster:
		for _, enemy := range l.enemies {
			enemy.SetNightMode(true)
		}
	}

	l.levelTiles[xTile][yTile] = Free
}

func (l *Level) UpdateEnemy(enemy *entities.Enemy) bool {
	if enemy.NightMode() && enemy.NightModeExpiredTime().Before(time.Now()) {
		enemy.SetNightMode(false)
	}

	rotation := enemy.GetDirection()
	nextX, nextY := enemy.CalculateNextPosition(rotation, base.GameScreenWidth, base.GameScreenHeight)

	// If a wall is encountered the coordinates do not change
	if l.CheckWallCollision(nextX, nextY) {
		enemy.SetStopped(true)
	} else {
		enemy.Move(rotation, base.GameScreenWidth, base.GameScreenWidth)
		enemy.SetStopped(false)
	}

	if l.checkHit(enemy.GetCoords()) {
		if !enemy.NightMode() {
			l.player.DecreaseHealth()
			return false
		}
		l.enemyKilled(enemy)
	}

	return true
}

func (l *Level) CheckWallCollision(x, y int) bool {
	xTileUp := int(math.Ceil(float64(x) / float64(l.tileSize)))
	yTileUp := int(math.Ceil(float64(y) / float64(l.tileSize)))

	xTileDown := int(math.Floor(float64(x) / float64(l.tileSize)))
	yTileDown := int(math.Floor(float64(y) / float64(l.tileSize)))

	return l.levelTiles[xTileUp%l.width][yTileUp%l.height] == Wall || l.levelTiles[xTileDown%l.width][yTileDown%l.height] == Wall ||
		l.levelTiles[xTileUp%l.width][yTileDown%l.height] == Wall || l.levelTiles[xTileDown%l.width][yTileUp%l.height] == Wall
}

func (l *Level) IsAllFoodEaten() bool {
	return l.foodCount == l.foodEaten
}

func (l *Level) GetEntityByCoordinates(c Coordinates) int {
	return l.levelTiles[c.X][c.Y]
}

func (l *Level) TileSize() int {
	return l.tileSize
}

func (l *Level) Width() int {
	return l.width
}

func (l *Level) Height() int {
	return l.height
}

func (l *Level) Score() int {
	return l.score
}

func (l *Level) Texts() []ScreenText {
	return l.texts
}

func (l *Level) Player() *entities.Pacman {
	return &l.player
}

func (l *Level) Enemies() []*entities.Enemy {
	return l.enemies
}

func (l *Level) createRandomPlayer(playerImage *ebiten.Image) entities.Pacman {
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Intn(l.width)
		y := rand.Intn(l.height)
		if l.levelTiles[x][y] == Free && !isEnemySpawn(Coordinates{x, y}) {
			p := entities.CreatePacman(x*l.tileSize, y*l.tileSize, playerImage)
			return p
		}
	}
}

func (l *Level) createFood() {
	for x := 0; x < l.width; x++ {
		for y := 0; y < l.height; y++ {
			if l.levelTiles[x][y] == Free && !isEnemySpawn(Coordinates{x, y}) {
				l.levelTiles[x][y] = Food
				l.foodCount++
			}
		}
	}
}

func (l *Level) createStrawberry() {
	rand.Seed(time.Now().UnixNano())
	countStrawberry := rand.Intn(strawberryMax-strawberryMin) + strawberryMin
	for countStrawberry > 0 {
		var x, y int
		for {
			x = rand.Intn(l.width)
			y = rand.Intn(l.height)
			if l.levelTiles[x][y] == Free && !isEnemySpawn(Coordinates{x, y}) {
				break
			}
		}
		l.levelTiles[x][y] = Strawberry
		countStrawberry--
	}
}
func (l *Level) createNightModeBoosters() {
	rand.Seed(time.Now().UnixNano())
	const boostersMax = 8
	const boostersMin = 2
	countBoosters := rand.Intn(boostersMax-boostersMin) + boostersMin
	for countBoosters > 0 {
		var x, y int
		for {
			x = rand.Intn(l.width)
			y = rand.Intn(l.height)
			if l.levelTiles[x][y] == Free && !isEnemySpawn(Coordinates{x, y}) {
				break
			}
		}
		l.levelTiles[x][y] = NightModeBooster
		countBoosters--
	}
}

func (l *Level) enemyKilled(enemy *entities.Enemy) {
	x, y := enemy.GetCoords()
	xTile := x / l.tileSize
	yTile := y / l.tileSize
	l.score += 400
	l.texts = append(l.texts, NewScreenText(xTile, yTile, "+400", time.Now().Add(textDuration*time.Second)))
	// Respawn enemy
	enemy.SetCoords(enemy.GetStartCoords())
	enemy.SetNightMode(false)
}

//nolint:gomnd
func (l *Level) checkHit(x, y int) bool {
	pacmanX, pacmanY := l.player.GetCoords()
	pacmanCenterX := (pacmanX + l.tileSize) / 2
	pacmanCenterY := (pacmanY + l.tileSize) / 2
	enemyCenterX := (x + l.tileSize) / 2
	enemyCenterY := (y + l.tileSize) / 2

	return math.Hypot(float64(enemyCenterX-pacmanCenterX), float64(enemyCenterY-pacmanCenterY)) < float64(l.tileSize)/3.0
}
