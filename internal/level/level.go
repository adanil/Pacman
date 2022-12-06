package level

import (
	"log"
	"math/rand"
	"pacman/internal/player"
	"time"
)

const (
	Free = iota
	Wall
	Player
	Bot
)

type Level struct {
	LevelTiles [][]int
	player     player.Player
}

func GenerateRandomLevel(widthTiles, heightTiles int) Level {
	level := Level{LevelTiles: make([][]int, widthTiles)}
	for x := 0; x < widthTiles; x++ {
		level.LevelTiles[x] = make([]int, heightTiles)
	}
	rand.Seed(time.Now().UnixNano())
	wallsCount := rand.Intn(widthTiles * heightTiles / 2)
	log.Println("Wall count:", wallsCount)
	for i := 0; i < wallsCount; i++ {
		x := rand.Intn(widthTiles)
		y := rand.Intn(heightTiles)
		if level.LevelTiles[x][y] == Free {
			level.LevelTiles[x][y] = Wall
		}
	}

	for {
		x := rand.Intn(widthTiles)
		y := rand.Intn(heightTiles)
		if level.LevelTiles[x][y] == Free {
			level.LevelTiles[x][y] = Player
			level.player.CreatePlayer(x, y)
			break
		}
	}
	return level

}
