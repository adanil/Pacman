package level

import (
	"log"
	"math/rand"
	"time"
)

type RandomLevelGenerator struct {
}

func (l *RandomLevelGenerator) CreateLevel(width, height, tileSize, enemyCount int) Level {
	level := Level{LevelTiles: make([][]int, width), Width: width, Height: height, TileSize: tileSize}
	for x := 0; x < width; x++ {
		level.LevelTiles[x] = make([]int, height)
	}
	rand.Seed(time.Now().UnixNano())
	wallsCount := rand.Intn(width * height / 2)
	log.Println("Wall count:", wallsCount)
	for i := 0; i < wallsCount; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		if level.LevelTiles[x][y] == Free {
			level.LevelTiles[x][y] = Wall
		}
	}
	return level
}
