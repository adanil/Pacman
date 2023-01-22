package level

import (
	"bufio"
	"os"
)

type ReadLevel struct {
	Filepath string
}

func (l *ReadLevel) CreateLevel(width, height, tileSize int) Level {
	level := Level{LevelTiles: make([][]int, width), Width: width, Height: height, TileSize: tileSize}
	for x := 0; x < width; x++ {
		level.LevelTiles[x] = make([]int, height)
	}

	file, _ := os.OpenFile(l.Filepath, os.O_RDONLY, 0666)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	h := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for index, c := range line {
			switch c {
			case '#':
				level.LevelTiles[index][h] = Wall
			}
		}
		h++
	}
	return level
}