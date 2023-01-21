package level

import (
	"math/rand"
	"time"
)

type RandomLevelGenerator struct {
}

func (l *RandomLevelGenerator) CreateLevel(width, height, tileSize int) Level {
	level := Level{LevelTiles: make([][]int, width), Width: width, Height: height, TileSize: tileSize}
	for x := 0; x < width; x++ {
		level.LevelTiles[x] = make([]int, height)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			level.LevelTiles[x][y] = Wall
		}
	}

	rand.Seed(time.Now().UnixNano())

	pathLen := rand.Intn(width*height/4) + (width*height)/4

	lenLine := 1
	direction := 0
	currX := rand.Intn(width)
	currY := rand.Intn(height)
	for pathLen != 0 {
		if isValid(level, currX, currY, width, height) {
			level.LevelTiles[currX][currY] = Free
			lenLine--
			pathLen--
		} else {
			lenLine = 0
		}

		if lenLine == 0 {
			direction = rand.Intn(4)
			lenLine = randIntWithRange(6, 10)
		}

		switch direction {
		case 0:
			currX++
		case 1:
			currX--
		case 2:
			currY++
		case 3:
			currY--
		}
		currX = (currX + width) % width
		currY = (currY + height) % height
	}
	return level
}

func isValid(level Level, x, y, width, height int) bool {
	if level.LevelTiles[x][y] == Free {
		return false
	}
	// a b c
	// d x g
	// m n k
	aX := coordDec(x, width)
	aY := coordDec(y, height)

	bX := x
	bY := coordDec(y, height)

	cX := coordInc(x, width)
	cY := coordDec(y, height)

	dX := coordDec(x, width)
	dY := y

	gX := coordInc(x, width)
	gY := y

	mX := coordDec(x, width)
	mY := coordInc(y, height)

	nX := x
	nY := coordInc(y, height)

	kX := coordInc(x, width)
	kY := coordInc(y, height)

	if (level.LevelTiles[aX][aY] == Free && level.LevelTiles[bX][bY] == Free && level.LevelTiles[dX][dY] == Free) ||
		(level.LevelTiles[dX][dY] == Free && level.LevelTiles[mX][mY] == Free && level.LevelTiles[nX][nY] == Free) ||
		(level.LevelTiles[bX][bY] == Free && level.LevelTiles[cX][cY] == Free && level.LevelTiles[gX][gY] == Free) ||
		(level.LevelTiles[nX][nY] == Free && level.LevelTiles[kX][kY] == Free && level.LevelTiles[gX][gY] == Free) {
		return false
	}
	return true

}

func randIntWithRange(from, to int) int {
	return rand.Intn(to-from) + from
}

func coordDec(val, module int) int {
	return (val - 1 + module) % module
}

func coordInc(val, module int) int {
	return (val + 1 + module) % module
}
