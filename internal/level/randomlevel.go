package level

import (
	"math/rand"
	"time"
)

type RandomLevelGenerator struct {
}

func (l *RandomLevelGenerator) CreateLevel(width, height, tileSize int) Level {
	level := NewLevel(width, height, tileSize)
	for x := 0; x < width; x++ {
		level.levelTiles[x] = make([]int, height)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			level.levelTiles[x][y] = Wall
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
			level.levelTiles[currX][currY] = Free
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
	if level.levelTiles[x][y] == Free {
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

	if (level.levelTiles[aX][aY] == Free && level.levelTiles[bX][bY] == Free && level.levelTiles[dX][dY] == Free) ||
		(level.levelTiles[dX][dY] == Free && level.levelTiles[mX][mY] == Free && level.levelTiles[nX][nY] == Free) ||
		(level.levelTiles[bX][bY] == Free && level.levelTiles[cX][cY] == Free && level.levelTiles[gX][gY] == Free) ||
		(level.levelTiles[nX][nY] == Free && level.levelTiles[kX][kY] == Free && level.levelTiles[gX][gY] == Free) {
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
