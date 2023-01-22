package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"pacman/internal/controllers"
	"pacman/internal/entities"
	"pacman/internal/level"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nfnt/resize"
)

/* TODO
*   Add menu
*	Add boosters
* 	Refactor code
*	Check code by linter
*	Optional: create AI generator for map
*	Optional: create AI controller for enemies
 */

const (
	frameModulo      = 150
	tileSize         = 32
	gameScreenWidth  = 928
	gameScreenHeight = 704
	widthTiles       = gameScreenWidth / tileSize
	heightTiles      = gameScreenHeight / tileSize
)

var (
	frameNumber        int
	gameLevel          level.Level
	keyboardController controllers.KeyboardHandler
	enemyController    controllers.EnemyController
	wallImage          *ebiten.Image //TODO delete this variable later
	foodImage          *ebiten.Image //TODO probably this too
	title              string
)

func init() {
	title = "Pacman"
	lvGenerator := level.Generator{Creator: &level.ReadLevel{Filepath: "maps/base"}}
	gameLevel = lvGenerator.CreateLevel(widthTiles, heightTiles, tileSize)

	readerWall, _ := os.Open("images/wall2.jpg")
	imgWall, _, _ := image.Decode(readerWall)
	resizedWallImage := resize.Resize(tileSize, tileSize, imgWall, resize.NearestNeighbor)
	wallImage = ebiten.NewImageFromImage(resizedWallImage)

	readerFood, _ := os.Open("images/dot.png")
	imgFood, _, _ := image.Decode(readerFood)
	resizedFoodImage := resize.Resize(tileSize/5, tileSize/5, imgFood, resize.NearestNeighbor)
	foodImage = ebiten.NewImageFromImage(resizedFoodImage)

	pacmanImage, _ := readImage("images/pacman-pack_v2/Pacmanx2.png")
	blueEnemyImage, _ := readImage("images/pacman-pack_v2/BlueEnemyx2.png")
	pinkEnemyImage, _ := readImage("images/pacman-pack_v2/PinkEnemyx2.png")
	redEnemyImage, _ := readImage("images/pacman-pack_v2/RedEnemyx2.png")

	pacman := CreateRandomPlayer(gameLevel, pacmanImage)
	blueEnemy := CreateRandomPlayer(gameLevel, blueEnemyImage)
	pinkEnemy := CreateRandomPlayer(gameLevel, pinkEnemyImage)
	redEnemy := CreateRandomPlayer(gameLevel, redEnemyImage)

	gameLevel.Player = pacman
	gameLevel.Enemies = append(gameLevel.Enemies, &blueEnemy, &pinkEnemy, &redEnemy)
	gameLevel.CreateFood()

	keyboardController = controllers.NewKeyboardHandler(&gameLevel)
	go keyboardController.HandlePressedButtons()
	enemyController = controllers.NewEnemyController(&gameLevel)

}

func readImage(imagePath string) (*ebiten.Image, error) {
	f, err := os.Open(imagePath)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

func CreateRandomPlayer(lv level.Level, playerImage *ebiten.Image) entities.Pacman {
	for {
		x := rand.Intn(lv.Width)
		y := rand.Intn(lv.Height)
		if lv.LevelTiles[x][y] == level.Free {
			p := entities.CreatePlayer(x, y, tileSize)
			p.Graphic = playerImage
			return p
		}
	}
}

type Game struct {
}

func (g *Game) Update() error {
	commands := keyboardController.GetKeyboardCommands()
	commands = append(commands, enemyController.GetCommands()...)
	for _, com := range commands {
		com.Execute()
	}
	gameLevel.UpdateAll()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	frameNumber = (frameNumber + 1) % frameModulo
	//TODO refactore duplicate code
	for x := 0; x < widthTiles; x++ {
		for y := 0; y < heightTiles; y++ {
			if gameLevel.LevelTiles[x][y] == level.Wall {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
				screen.DrawImage(wallImage, op)
			} else if gameLevel.LevelTiles[x][y] == level.Food {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize)+tileSize/2, float64(y*tileSize)+tileSize/2)
				screen.DrawImage(foodImage, op)
			}
		}
	}

	g.drawPacman(screen)
	g.drawEnemies(screen)
	g.drawTitle(screen)
	g.drawScore(screen)

}

func (g *Game) drawTitle(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(gameScreenWidth/2-20, 20)
	text.DrawWithOptions(screen, title, basicfont.Face7x13, op)
}

func (g *Game) drawScore(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(35, gameScreenHeight-15)
	text.DrawWithOptions(screen, "Score: "+strconv.Itoa(gameLevel.Player.Score), basicfont.Face7x13, op)
}

func (g *Game) drawMap(screen *ebiten.Image) {
	for x := 0; x < widthTiles; x++ {
		for y := 0; y < heightTiles; y++ {
			if gameLevel.LevelTiles[x][y] == level.Wall {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
				screen.DrawImage(wallImage, op)
			}
		}
	}
}

func (g *Game) drawEnemies(screen *ebiten.Image) {
	state := frameNumber / (frameModulo / 2)
	for _, enemy := range gameLevel.Enemies {
		x, y := enemy.GetCoords()
		dir := enemy.GetDirection()
		fmt.Printf("Coords: %v, %v ,%v\n", x, y, dir)
		op := &ebiten.DrawImageOptions{}
		px, py := enemy.GetCoords()
		op.GeoM.Translate(float64(px), float64(py))
		enemyImage := enemy.GetGraphic().SubImage(image.Rect(6+40*(enemy.GetDirection()*2+state), 0, 6+40*(enemy.GetDirection()*2+state)+30, 38)).(*ebiten.Image)
		screen.DrawImage(enemyImage, op)
	}
}

func (g *Game) drawPacman(screen *ebiten.Image) {
	state := frameNumber / (frameModulo / 3)
	pacman := gameLevel.Player
	op := &ebiten.DrawImageOptions{}
	px, py := pacman.GetCoords()
	op.GeoM.Translate(float64(px), float64(py))
	pacmanImage := (*ebiten.Image)(nil)
	if state == 2 {
		pacmanImage = pacman.GetGraphic().SubImage(image.Rect(296, 0, 324, 30)).(*ebiten.Image)
	} else {
		state = (state + 1) % 2
		pacmanImage = pacman.GetGraphic().SubImage(image.Rect(38*(pacman.GetDirection()*2+state), 0, 38*(pacman.GetDirection()*2+state)+30, 30)).(*ebiten.Image)
	}
	screen.DrawImage(pacmanImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameScreenWidth, gameScreenHeight
}

func main() {
	ebiten.SetWindowSize(gameScreenWidth, gameScreenHeight)
	ebiten.SetWindowTitle("Pacman")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
