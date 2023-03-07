package level

import (
	"bufio"
	"os"
)

type Creator interface {
	CreateLevel(width, height, tileSIze int) Level
}

type ReadLevel struct {
	Filepath string
}

func (l *ReadLevel) CreateLevel(width, height, tileSize int) Level {
	level := NewLevel(width, height, tileSize)
	for x := 0; x < width; x++ {
		level.levelTiles[x] = make([]TileObject, height)
	}

	file, _ := os.OpenFile(l.Filepath, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	h := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for index, c := range line {
			if c == '#' {
				level.levelTiles[index][h] = Wall
			}
		}
		h++
	}
	return level
}
