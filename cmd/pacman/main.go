package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"pacman/internal/level"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nfnt/resize"
)

const (
	tileSize         = 30
	gameScreenWidth  = 640
	gameScreenHeight = 480
	widthTiles       = gameScreenWidth / tileSize
	heightTiles      = gameScreenHeight / tileSize
)

var (
	wallImage   *ebiten.Image
	pacmanImage *ebiten.Image
	gameLevel   level.Level
)

func Init() {
	readerWall, err := os.Open("images/wall.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer readerWall.Close()
	readerPacman, err := os.Open("images/pacman.png")
	if err != nil {
		log.Fatal(err)
	}

	imgWall, _, err := image.Decode(readerWall)
	if err != nil {
		log.Fatal(err)
	}
	imgPacman, _, err := image.Decode(readerPacman)
	if err != nil {
		log.Fatal(err)
	}
	resizedWallImage := resize.Resize(tileSize, tileSize, imgWall, resize.NearestNeighbor)
	wallImage = ebiten.NewImageFromImage(resizedWallImage)

	resizedPacmanImage := resize.Resize(tileSize, tileSize, imgPacman, resize.NearestNeighbor)
	pacmanImage = ebiten.NewImageFromImage(resizedPacmanImage)

}

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < widthTiles; x++ {
		for y := 0; y < heightTiles; y++ {
			if gameLevel.LevelTiles[x][y] == level.Wall {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
				screen.DrawImage(wallImage, op)
			} else if gameLevel.LevelTiles[x][y] == level.Player {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
				screen.DrawImage(pacmanImage, op)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameScreenWidth, gameScreenHeight
}

func main() {
	Init()
	gameLevel = level.GenerateRandomLevel(widthTiles, heightTiles)
	ebiten.SetWindowSize(gameScreenWidth, gameScreenHeight)
	ebiten.SetWindowTitle("Pacman")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
