package base

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"pacman/internal/controllers"
	"pacman/internal/entities"
	"pacman/internal/level"
)

const (
	frameModulo      = 150
	tileSize         = 32
	gameScreenWidth  = 928
	gameScreenHeight = 704
	widthTiles       = gameScreenWidth / tileSize
	heightTiles      = gameScreenHeight / tileSize
	defaultDPI       = 64
)

var (
	frameNumber        int
	gameLevel          level.Level
	keyboardController controllers.KeyboardHandler
	enemyController    controllers.EnemyController
	wallImage          *ebiten.Image //TODO delete this variable later
	foodImage          *ebiten.Image //TODO probably this too
	title              string
	baseFont           *opentype.Font
	yellowColor        color.RGBA
	foodCount          int
)

// TODO Release this code after restart
func init() {
	fontReader, err := os.ReadFile("fonts/quinquefive/QuinqueFive.ttf")
	baseFont, err = opentype.Parse(fontReader)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	yellowColor = color.RGBA{
		R: 230,
		G: 245,
		B: 39,
		A: 255,
	}

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
	foodCount = gameLevel.CreateFood()

	keyboardController = controllers.NewKeyboardHandler(&gameLevel)
	go keyboardController.HandlePressedButtons()
	enemyController = controllers.NewEnemyController(&gameLevel)

}

func GetFont(baseFont *opentype.Font, size, dpi float64) (font.Face, error) {
	return opentype.NewFace(baseFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
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
	currentState State
}

func (g *Game) SetState(state State) {
	g.currentState = state
}

func (g *Game) Update() error {
	return g.currentState.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentState.Draw(screen)
}

// TODO CHECK params
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameScreenWidth, gameScreenHeight
}
